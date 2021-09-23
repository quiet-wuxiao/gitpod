{
  prometheusAlerts+:: {
    groups+: [
      {
        name: 'gitpod-component-ws-daemon-alerts',
        rules: [
          {
            alert: 'GitpodWsDaemonCrashLooping',
            labels: {
              severity: 'critical',
            },
            annotations: {
              runbook_url: 'https://github.com/gitpod-io/observability/blob/main/runbooks/GitpodWsDaemonCrashLooping.md',
              summary: 'Ws-daemon is crashlooping.',
              description: 'Pod {{ $labels.namespace }}/{{ $labels.pod }} ({{ $labels.container }}) is restarting {{ printf "%.2f" $value }} times / 10 minutes.',
            },
            expr: |||
              increase(kube_pod_container_status_restarts_total{container="ws-daemon"}[10m]) > 0
            |||,
          },
        ],
      },
    ],
  },
}