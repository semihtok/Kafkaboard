(function($) {
  "use strict";
  $(function() {
    if ($("#realtimeChart").length) {
      Highcharts.setOptions({
        global: {
          useUTC: false
        }
      });

      var lastItem = [];
      var incomeItemCount = 0;
      var totalItemCount = 0;
      var index = 0;

      Highcharts.chart("realtimeChart", {
        chart: {
          type: "spline",
          animation: Highcharts.svg, // don't animate in old IE
          marginRight: 10,
          events: {
            load: function() {
              // set up the updating of the chart each second
              var series = this.series[0];

              var kafkaRequest = {
                name: $("#txtTopic").val()
              };

              setInterval(function() {
                kafkaRequest.name = $("#txtTopic").val();

                if (kafkaRequest.name != "") {
                  $.ajax({
                    url: "/kafka",
                    type: "post",
                    data: JSON.stringify(kafkaRequest),
                    success: function(response) {
                      if (index != 0) {
                        var r = response.result;

                        if (r.length > 0 && lastItem.equals(r) == false) {
                          if (incomeItemCount > 0) {
                            incomeItemCount = totalItemCount - r.length;
                          } else {
                            incomeItemCount = r.length;
                          }
                          $(".latestMessages tr:last").after(
                            '<tr> <td class="text-left">' +
                              GetCurrentTime() +
                              '</td><td><div class="flag"><i class="flag-icon flag-icon-tr"></i></div></td><td class="text-right">' +
                              r +
                              "</td></tr>"
                          );
                          lastItem = r;
                        } else {
                          incomeItemCount = 0;
                        }
                      } else {
                        index++;
                        totalItemCount = response.result.length;
                      }
                    },
                    error: function(jqXHR, textStatus, errorThrown) {
                      console.log(textStatus, errorThrown);
                    }
                  });
                } else {
                  var x = new Date().getTime(), // current time
                    y = 0;
                  series.addPoint([x, y], true, true);
                }

                var x = new Date().getTime(), // current time
                  y = incomeItemCount;
                series.addPoint([x, y], true, true);
              }, 2000);
            }
          }
        },
        title: {
          text: "Real Time Topics"
        },
        xAxis: {
          type: "datetime",
          tickPixelInterval: 150
        },
        yAxis: {
          title: {
            text: "Value"
          },
          plotLines: [
            {
              value: 0,
              width: 1,
              color: "#808080"
            }
          ]
        },
        tooltip: {
          formatter: function() {
            return (
              "<b>" +
              this.series.name +
              "</b><br/>" +
              Highcharts.dateFormat("%Y-%m-%d %H:%M:%S", this.x) +
              "<br/>" +
              Highcharts.numberFormat(this.y, 2)
            );
          }
        },
        legend: {
          enabled: false
        },
        exporting: {
          enabled: false
        },
        series: [
          {
            name: "",
            color: "#8a8be2",
            data: (function() {
              // generate an array of random data
              var data = [],
                time = new Date().getTime(),
                i;

              for (i = -2; i <= 0; i += 1) {
                data.push({
                  x: time + i * 1000,
                  y: 0
                });
              }
              return data;
            })()
          }
        ]
      });
    }

    $("#btnSendTopic").click(function() {
      var topicName = $("#txtTopic").val();
      var message = $("#txtMessage").val();

      var topicModel = {
        Name: topicName,
        Message: message
      };

      $.ajax({
        url: "/actions",
        type: "post",
        data: JSON.stringify(topicModel),
        success: function(response) {
          if (response.status === "OK") {
            ohSnap("Message succesfully sent", {
              duration: "5000"
            });
          } else {
            ohSnap("Something happened and message couldn't send", {
              color: "red",
              duration: "5000"
            });
          }
        },
        error: function(jqXHR, textStatus, errorThrown) {
          console.log(textStatus, errorThrown);
        }
      });
    });
  });
})(jQuery);

if (Array.prototype.equals)
  console.warn(
    "Overriding existing Array.prototype.equals. Possible causes: New API defines the method, there's a framework conflict or you've got double inclusions in your code."
  );
// attach the .equals method to Array's prototype to call it on any array
Array.prototype.equals = function(array) {
  // if the other array is a falsy value, return
  if (!array) return false;

  // compare lengths - can save a lot of time
  if (this.length != array.length) return false;

  for (var i = 0, l = this.length; i < l; i++) {
    // Check if we have nested arrays
    if (this[i] instanceof Array && array[i] instanceof Array) {
      // recurse into the nested arrays
      if (!this[i].equals(array[i])) return false;
    } else if (this[i] != array[i]) {
      // Warning - two different object instances will never be equal: {x:20} != {x:20}
      return false;
    }
  }
  return true;
};
// Hide method from for-in loops
Object.defineProperty(Array.prototype, "equals", { enumerable: false });

Number.prototype.padLeft = function(base, chr) {
  var len = String(base || 10).length - String(this).length + 1;
  return len > 0 ? new Array(len).join(chr || "0") + this : this;
};

function GetCurrentTime() {
  var date = new Date(),
    year = date.getFullYear(),
    month = (date.getMonth() + 1).toString(),
    formatedMonth = month.length === 1 ? "0" + month : month,
    day = date.getDate().toString(),
    formatedDay = day.length === 1 ? "0" + day : day,
    hour = date.getHours().toString(),
    formatedHour = hour.length === 1 ? "0" + hour : hour,
    minute = date.getMinutes().toString(),
    formatedMinute = minute.length === 1 ? "0" + minute : minute,
    second = date.getSeconds().toString(),
    formatedSecond = second.length === 1 ? "0" + second : second;
  return (
    formatedDay +
    "-" +
    formatedMonth +
    "-" +
    year +
    " " +
    formatedHour +
    ":" +
    formatedMinute +
    ":" +
    formatedSecond
  );
}
