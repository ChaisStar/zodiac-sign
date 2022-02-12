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

var ZODIAC_TYPES = [
  "Traditional",
  "Chinese",
  "Common",
  "Detailed",
  "Love"
];

var app = new Vue({
  el: '#app',
  components: {
    vuejsDatepicker
  },
  data: {
    signs: ZODIAC_SIGNS,
    types: ZODIAC_TYPES,
    startDate: new Date(),
    endDate: new Date(),
    selectedSigns: ZODIAC_SIGNS.map((_, i) => i + 1),
    selectedTypes: ZODIAC_TYPES.map((_, i) => i),
  },
  methods: {
    download: function () {
      axios({
        method: 'post',
        url: '/api',
        data: {
          signs: this.selectedSigns,
          startDate: this.startDate.toISOString().split('T')[0],
          endDate: this.endDate.toISOString().split('T')[0],
          types: this.selectedTypes
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
              a.download = `horoscope_${this.startDate.toISOString().split('T')[0]}_${this.endDate.toISOString().split('T')[0]}`;
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