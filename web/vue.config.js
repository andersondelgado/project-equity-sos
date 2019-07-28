const ForkTsCheckerWebpackPlugin = require('fork-ts-checker-webpack-plugin');
const os=require('os');
module.exports = {
  pluginOptions: {
    i18n: {
      locale: 'es',
      fallbackLocale: 'en',
      localeDir: 'locales',
      enableInSFC: true
    }
  },
  configureWebpack: config => {

    // get a reference to the existing ForkTsCheckerWebpackPlugin
    const existingForkTsChecker = config.plugins.filter(
      p => p instanceof ForkTsCheckerWebpackPlugin,
    )[0];

    // remove the existing ForkTsCheckerWebpackPlugin
    // so that we can replace it with our modified version
    config.plugins = config.plugins.filter(
      p => !(p instanceof ForkTsCheckerWebpackPlugin),
    );

    // copy the options from the original ForkTsCheckerWebpackPlugin
    // instance and add the memoryLimit property
    const forkTsCheckerOptions = existingForkTsChecker.options;
    forkTsCheckerOptions.memoryLimit = 8192;

    config.plugins.push(new ForkTsCheckerWebpackPlugin(forkTsCheckerOptions));
  }
  // chainWebpack: config => {
  //       config
  //           .plugin('fork-ts-checker')
  //           .tap(args => {
  //               let totalmem=Math.floor(os.totalmem()/1024/1024); //get OS mem size
  //               let allowUseMem= totalmem>2500? 2048:1000;
  //               args[0].memoryLimit = allowUseMem;
  //               return args
  //           })
  //   }
}
