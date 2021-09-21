package io.gitpod.gitpodprotocol.api;

import java.util.concurrent.CompletableFuture;

import org.eclipse.lsp4j.jsonrpc.services.JsonRequest;

import io.gitpod.gitpodprotocol.api.entities.SendHeartBeatOptions;
import io.gitpod.gitpodprotocol.api.entities.User;

public interface GitpodServer {
    @JsonRequest
    CompletableFuture<User> getLoggedInUser();

    @JsonRequest
    CompletableFuture<Void> sendHeartBeat(SendHeartBeatOptions options);
}
