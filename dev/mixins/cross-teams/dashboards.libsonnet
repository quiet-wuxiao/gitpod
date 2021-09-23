(import './dashboards/gitpod-overview.libsonnet') +
(import './dashboards/gitpod-nodes-overview.libsonnet') +
(import './dashboards/gitpod-admin-workspace.libsonnet') +
(import './dashboards/gitpod-admin-node.libsonnet') +
(import './dashboards/SLOs/workspace-startup-time.libsonnet') +
{
    grafanaDashboards+:: {
        // Import raw json files here.
        // Example:
        // 'my-new-dashboard.json': (import 'dashboards/components/new-component.json'),
    },
}