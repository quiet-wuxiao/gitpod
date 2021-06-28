/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { injectable, inject } from "inversify";
import { UserDeletionService } from "../../../src/user/user-deletion-service";
import { SubscriptionService } from "@gitpod/gitpod-payment-endpoint/lib/accounting/subscription-service";
import { Plans } from "@gitpod/gitpod-protocol/lib/plans";
import { ChargebeeService } from "./chargebee-service";
import { EnvEE } from "../env";
import { IAnalyticsWriter } from '@gitpod/gitpod-protocol/lib/util/analytics';
@injectable()
export class UserDeletionServiceEE extends UserDeletionService {
    @inject(ChargebeeService) protected readonly chargebeeService: ChargebeeService;
    @inject(SubscriptionService) protected readonly subscriptionService: SubscriptionService;
    @inject(EnvEE) protected readonly env: EnvEE;
    @inject(IAnalyticsWriter) protected readonly analytics: IAnalyticsWriter;

    async deleteUser(id: string): Promise<void> {
        const user = await this.db.findUserById(id);
        if (!user) {
            throw new Error(`No user with id ${id} found!`);
        }

        const now = new Date().toISOString();
        if (this.env.enablePayment) {
            const subscriptions = await this.subscriptionService.getNotYetCancelledSubscriptions(user, now);
            for (const subscription of subscriptions) {
                const planId = subscription.planId!;
                if (Plans.isFreeNonTransientPlan(planId)) {
                    // only delete those plans that are persisted in the DB
                    await this.subscriptionService.unsubscribe(user.id, now, planId);
                } else if (Plans.isFreePlan(planId)) {
                    // we do not care about transient plans
                    continue;
                } else {
                    // cancel Chargebee subscriptions
                    const subscriptionId = subscription.uid;
                    const chargebeeSubscriptionId = subscription.paymentReference!;
                    await this.chargebeeService.cancelSubscription(chargebeeSubscriptionId, { userId: user.id }, { subscriptionId, chargebeeSubscriptionId });
                }
            }
        }

        // track the deletion event in downstream analytics tools
        this.analytics.track({
            userId: user.id,
            event: "deletion",
            properties: {
                "deleted_at":now
            }
        });

        return super.deleteUser(id);
    }
}