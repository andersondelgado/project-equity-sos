$(document).ready(function() {
    md.initFormExtendedDatetimepickers();
    if ($('.slider').length != 0) {
        md.initSliders();
    }

    md.initDashboardPageCharts();
    md.initVectorMap();
});