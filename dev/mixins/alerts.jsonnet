std.manifestYamlDoc(
    (import './cross-teams/mixin.libsonnet').prometheusAlerts +
    (import './IDE/mixin.libsonnet').prometheusAlerts +
    (import './meta/mixin.libsonnet').prometheusAlerts +
    (import './workspace/mixin.libsonnet').prometheusAlerts
)