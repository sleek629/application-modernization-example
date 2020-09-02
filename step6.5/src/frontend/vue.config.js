 module.exports = {
   devServer: {
     proxy: {
       "/WordAPI": {
         target: "http://envoy:8001"
       }
     }
   }
 };
