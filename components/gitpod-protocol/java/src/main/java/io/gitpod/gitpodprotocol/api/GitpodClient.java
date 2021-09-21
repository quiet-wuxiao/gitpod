package io.gitpod.gitpodprotocol.api;

public interface GitpodClient {
    void connect(GitpodServer server);

    GitpodServer server();
}
