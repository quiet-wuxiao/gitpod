{
  _config+:: {
    // Make it possible to generate dashboards compatible with multicluster installations
    showMultiCluster: true,
    clusterLabel: 'cluster',

    gitpodURL: 'https://gitpod.io',

    dashboardNamePrefix: 'Gitpod / ',
    dashboardTags: ['gitpod-mixin'],
  },
}
