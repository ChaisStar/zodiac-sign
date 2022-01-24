var ZODIAC_SIGNS = [
  "Aries",
  "Taurus",
  "Gemini",
  "Cancer",
  "Leo",
  "Virgo",
  "Libra",
  "Scorpio",
  "Sagittarius",
  "Capricorn",
  "Aquarius",
  "Pisces"
];

var app = new Vue({
  el: '#app',
  components: {
    vuejsDatepicker
  },
  data: {
    signs: ZODIAC_SIGNS,
    startDate: new Date(),
    endDate: new Date(),
    selectedSigns: ZODIAC_SIGNS.map((_, i) => i + 1)
  },
  methods: {
    download: function () {
      axios({
        method: 'post',
        url: '/api',
        data: {
          signs: this.selectedSigns,
          startDate: this.startDate.toISOString().split('T')[0],
          endDate: this.endDate.toISOString().split('T')[0]
        },
        responseType: 'blob'
      })
      .then(response => {
            let blob = new Blob([response.data], { type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" }),
              downloadUrl = window.URL.createObjectURL(blob),
              filename = "",
              disposition = response.headers["content-disposition"];
  
            if (disposition && disposition.indexOf("attachment") !== -1) {
              let filenameRegex = /filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/,
                matches = filenameRegex.exec(disposition);
  
              if (matches != null && matches[1]) {
                filename = matches[1].replace(/['"]/g, "");
              }
            }
  
            let a = document.createElement("a");
            if (typeof a.download === "undefined") {
              window.location.href = downloadUrl;
            } else {
              a.href = downloadUrl;
              a.download = "horoscope";
              document.body.appendChild(a);
              a.click();
            }
          });
      // axios
      //   .post(
      //     '/api',
      //     {
      //       signs: this.selectedSigns,
      //       startDate: this.startDate.toISOString().split('T')[0],
      //       endDate: this.endDate.toISOString().split('T')[0]
      //     },
      //     "blob")
      //   .then(response => {
      //     let blob = new Blob([response.data], { type: "application/zip" }),
      //       downloadUrl = window.URL.createObjectURL(blob),
      //       filename = "",
      //       disposition = response.headers["content-disposition"];

      //     if (disposition && disposition.indexOf("attachment") !== -1) {
      //       let filenameRegex = /filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/,
      //         matches = filenameRegex.exec(disposition);

      //       if (matches != null && matches[1]) {
      //         filename = matches[1].replace(/['"]/g, "");
      //       }
      //     }

      //     let a = document.createElement("a");
      //     if (typeof a.download === "undefined") {
      //       window.location.href = downloadUrl;
      //     } else {
      //       a.href = downloadUrl;
      //       a.download = filename;
      //       document.body.appendChild(a);
      //       a.click();
      //     }
      //   })
    }
  }
})