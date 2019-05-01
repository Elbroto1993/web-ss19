"use strict";

// Get kastenid from url
let urlParam = function(name, w) {
  w = w || window;
  var rx = new RegExp("[&|?]" + name + "=([^&#]+)"),
    val = w.location.search.match(rx);
  return !val ? "" : val[1];
};
let id = urlParam("Use_Id");

document.addEventListener("DOMContentLoaded", displayAllData);

// Get data for karteikasten sidebar counter
function displayAllData() {
  // Get kasten counter
  // Prepare http request
  let xhr3 = new XMLHttpRequest();
  xhr3.responseType = `json`;
  let url = "http://localhost:8080/karteikasten/";
  xhr3.open("GET", url, true);
  xhr3.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr3.onreadystatechange = function() {
    if (xhr3.readyState == 4 && xhr3.status == 200) {
      // Count all public Karteikaesten
      let counterPublic = 0;
      for (let i = 0; i < xhr3.response.length; i++) {
        if (xhr3.response[i].private == "false") {
          counterPublic++;
        }
      }
      document.getElementById(
        "indexKastenCounterPublic"
      ).innerHTML = counterPublic;
    }
  };
  // GET data
  xhr3.send(null);

  // Get user karten counter
  let indexKartenCounterUser = document.getElementById(
    "indexKartenCounterUser"
  );
  // Prepare http request
  let xhr4 = new XMLHttpRequest();
  xhr4.responseType = `json`;
  url = "http://localhost:8080/karteikasten/user";
  xhr4.open("GET", url, true);
  xhr4.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr4.onreadystatechange = function() {
    if (xhr4.readyState == 4 && xhr4.status == 200) {
      if (xhr4.response.length != null) {
        indexKartenCounterUser.innerHTML = xhr4.response.length;
      }
    }
  };
  // GET data
  xhr4.send(null);
  // Get all infos about user
  let xhr6 = new XMLHttpRequest();
  url = "http://localhost:8080/user/id";
  xhr6.open("GET", url, true);
  //xhr6.responseType = `json`;
  xhr6.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr6.onreadystatechange = function() {
    if (xhr6.readyState == 4 && xhr6.status == 200) {
      let res = JSON.parse(xhr6.response);
      document.getElementById("logoutUsername").innerHTML = res.username;
    }
  };
  // GET data
  xhr6.send(null);
  // Get all karten pro kasten
  let xhr2 = new XMLHttpRequest();
  let url3 = `http://localhost:8080/karteikarte/kasten/${id}`;
  xhr2.responseType = `json`;
  xhr2.open("GET", url3, true);
  xhr2.onreadystatechange = function() {
    if (xhr2.readyState == 4 && xhr2.status == 200) {
      let data = xhr2.response;
      // Only do the following if kasten contains karten
      if (data) {
        // Fill element (anzahl karten) with data from length of slice, if data isn't null
        if (!data) {
          document.getElementById("viewCountKarten").innerHTML = 0;
        } else {
          document.getElementById("viewCountKarten").innerHTML = data.length;
        }
        // Loop through all karten and count how many karten are in each fach
        let fach0 = 0;
        let fach1 = 0;
        let fach2 = 0;
        let fach3 = 0;
        let fach4 = 0;
        for (let i = 0; i < data.length; i++) {
          switch (data[i].fach) {
            case "0":
              fach0++;
              break;
            case "1":
              fach1++;
              break;
            case "2":
              fach2++;
              break;
            case "3":
              fach3++;
              break;
            case "4":
              fach4++;
              break;
          }
        }
        // Calculate progress for kasten
        let progressHelp1 = 0;
        progressHelp1 += 1 * fach1;
        progressHelp1 += 2 * fach2;
        progressHelp1 += 3 * fach3;
        progressHelp1 += 4 * fach4;
        let progress = Math.floor((progressHelp1 / (4 * data.length)) * 100);
        // Update kasten to add progress
        let xhr5 = new XMLHttpRequest();
        let url4 = `http://localhost:8080/karteikasten/kasten/${id}`;
        xhr5.responseType = `json`;
        xhr5.open("GET", url4, true);
        xhr5.onreadystatechange = function() {
          if (xhr5.readyState == 4 && xhr5.status) {
            let xhr6 = new XMLHttpRequest();
            let url5 = `http://localhost:8080/karteikasten/update`;
            xhr6.responseType = `json`;
            xhr6.open("POST", url5, true);
            xhr6.onreadystatechange = function() {
              if (xhr6.readyState == 4 && xhr6.status) {
              }
            };
            let data = JSON.stringify({
              kastenid: xhr5.response.kastenid,
              titel: xhr5.response.titel,
              beschreibung: xhr5.response.beschreibung,
              private: xhr5.response.private,
              kategorie: xhr5.response.kategorie,
              ueberkategorie: xhr5.response.ueberkategorie,
              fortschritt: progress.toString()
            });
            xhr6.send(data);
          }
        };
        xhr5.send(null);

        document.getElementById("fach0").innerHTML = fach0 + " ";
        document.getElementById("fach1").innerHTML = fach1 + " ";
        document.getElementById("fach2").innerHTML = fach2 + " ";
        document.getElementById("fach3").innerHTML = fach3 + " ";
        document.getElementById("fach4").innerHTML = fach4 + " ";

        // Show random karte from random fach from kasten
        let randomFach;
        let containsKarten = false;
        // Check if there are karten are in the random fach
        while (!containsKarten) {
          randomFach = zufallsFach();
          switch (randomFach) {
            case 0:
              if (fach0 != 0) containsKarten = true;
              break;
            case 1:
              if (fach1 != 0) containsKarten = true;
              break;
            case 2:
              if (fach2 != 0) containsKarten = true;
              break;
            case 3:
              if (fach3 != 0) containsKarten = true;
              break;
            case 4:
              if (fach4 != 0) containsKarten = true;
              break;
          }
        }
        // Loop through every karte until one with the fach is found
        for (let i = 0; i < data.length; i++) {
          if (data[i].fach == randomFach) {
            let divToAdd = `
            <div class="viewGrid">
            <div class="viewPositionTitle">Titel</div>
            <div class="flexNoWrap viewPositionTitle">
              <div><strong>${data[i].titel}</strong></div>
              <div id="singleCardDiv">
                <div class="flexNoWrap">
                  <div>
                    <div class="hexagon hexagon2 viewHexagon">
                      <div class="hexagon-in1">
                        <div class="hexagon-in2"></div>
                      </div>
                    </div>
                    <div class="viewFachInsideHexagon">
                      <p>0</p>
                    </div>
                  </div>
                  <div>
                    <div class="hexagon hexagon2 viewHexagon">
                      <div class="hexagon-in1">
                        <div class="hexagon-in2"></div>
                      </div>
                    </div>
                    <div class="viewFachInsideHexagon">
                      <p>1</p>
                    </div>
                  </div>
                  <div>
                    <div class="hexagon hexagon2 viewHexagon">
                      <div class="hexagon-in1">
                        <div class="hexagon-in2"></div>
                      </div>
                    </div>
                    <div class="viewFachInsideHexagon">
                      <p>2</p>
                    </div>
                  </div>
                  <div>
                    <div class="hexagon hexagon2 viewHexagon">
                      <div class="hexagon-in1">
                        <div class="hexagon-in2"></div>
                      </div>
                    </div>
                    <div class="viewFachInsideHexagon">
                      <p>3</p>
                    </div>
                  </div>
                  <div>
                    <div class="hexagon hexagon2 viewHexagon">
                      <div class="hexagon-in1">
                        <div class="hexagon-in2"></div>
                      </div>
                    </div>
                    <div class="viewFachInsideHexagon">
                      <p>4</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div>Frage</div>
              <div>
                ${data[i].frage}
              </div>
              <div>Antwort</div>
              <div></div>
            </div>
            <div class="lernCenterButton" id="${data[i].karteid}">
              <button class="btnYellow aufdeckenButton" id="aufdeckenButton">Aufdecken</button>
            </div>
            <div class="lernRightButton">
              <button class="btnYellow" id="skipButton">Überspringen</button>
            </div>
            `;
            document.getElementById("lernShowKarte").innerHTML = divToAdd;
            // Add green background to hexagon fach
            let hexToAddOne = document
              .getElementById("singleCardDiv")
              .getElementsByClassName("viewHexagon");
            let hexToAddTwo = document
              .getElementById("singleCardDiv")
              .getElementsByClassName("hexagon-in2");
            for (let j = 0; j < 5; j++) {
              if (data[i].fach == j) {
                hexToAddOne[j].id = "hex1";
                hexToAddTwo[j].id = "hex2";
              }
            }
            // Add event listener to buttons
            document
              .getElementById("aufdeckenButton")
              .addEventListener("click", karteAufdecken);
            document
              .getElementById("skipButton")
              .addEventListener("click", karteSkippen);
            break;
          }
        }
      }
    }
  };
  xhr2.send(null);
  // Get all infos about kasten
  let xhr = new XMLHttpRequest();
  url = `http://localhost:8080/karteikasten/kasten/${id}`;
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let data = xhr.response;

      // Fill elements (name, kategorie, fortschritt) with data from kasten
      document.getElementById("viewName").innerHTML = data.titel;
      document.getElementById("viewKategorie").innerHTML =
        data.ueberkategorie + ">" + data.kategorie;
      document.getElementById("viewFortschritt").innerHTML =
        data.fortschritt + " %";
      // Get username for createdById
      let xhr2 = new XMLHttpRequest();
      let url2 = `http://localhost:8080/user/id/${data.createdByUserId}`;
      xhr2.responseType = `json`;
      xhr2.open("GET", url2, true);
      xhr2.onreadystatechange = function() {
        if (xhr2.readyState == 4 && xhr2.status == 200) {
          let data = xhr2.response;
          // Fill element (erstellt von) with data from createdByUserId
          document.getElementById("viewCreatedBy").innerHTML = data.username;
        }
      };
      xhr2.send(null);
    }
  };
  xhr.send(null);
}
// Functions for karten buttons
function karteAufdecken() {
  let id = this.parentNode.id;
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/karteikarte/${id}`;
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let karte = xhr.response;
      let divToAdd = `
      <div class="viewGrid">
      <div class="viewPositionTitle">Titel</div>
      <div class="flexNoWrap viewPositionTitle">
        <div><strong>${karte.titel}</strong></div>
        <div id="singleCardDiv">
          <div class="flexNoWrap">
            <div>
              <div class="hexagon hexagon2 viewHexagon">
                <div class="hexagon-in1">
                  <div class="hexagon-in2"></div>
                </div>
              </div>
              <div class="viewFachInsideHexagon">
                <p>0</p>
              </div>
            </div>
            <div>
              <div class="hexagon hexagon2 viewHexagon">
                <div class="hexagon-in1">
                  <div class="hexagon-in2"></div>
                </div>
              </div>
              <div class="viewFachInsideHexagon">
                <p>1</p>
              </div>
            </div>
            <div>
              <div class="hexagon hexagon2 viewHexagon">
                <div class="hexagon-in1">
                  <div class="hexagon-in2"></div>
                </div>
              </div>
              <div class="viewFachInsideHexagon">
                <p>2</p>
              </div>
            </div>
            <div>
              <div class="hexagon hexagon2 viewHexagon">
                <div class="hexagon-in1">
                  <div class="hexagon-in2"></div>
                </div>
              </div>
              <div class="viewFachInsideHexagon">
                <p>3</p>
              </div>
            </div>
            <div>
              <div class="hexagon hexagon2 viewHexagon">
                <div class="hexagon-in1">
                  <div class="hexagon-in2"></div>
                </div>
              </div>
              <div class="viewFachInsideHexagon">
                <p>4</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div>Frage</div>
      <div>
        ${karte.frage}
      </div>
      <div>Antwort</div>
      <div>
        ${karte.antwort}
      </div>
      </div>
      <div class="lernCenterButton" id="${id}">
        <button class="btnGreen" id="richtigButton">
         Richtig
       </button>
        <button class="btnRed" id="falschButton">
         Falsch
       </button>
     </div>
      `;
      document.getElementById("lernShowKarte").innerHTML = divToAdd;
      // Add green background to hexagon fach
      let hexToAddOne = document
        .getElementById("singleCardDiv")
        .getElementsByClassName("viewHexagon");
      let hexToAddTwo = document
        .getElementById("singleCardDiv")
        .getElementsByClassName("hexagon-in2");
      for (let j = 0; j < 5; j++) {
        if (karte.fach == j) {
          hexToAddOne[j].id = "hex1";
          hexToAddTwo[j].id = "hex2";
        }
      }
      // Add event listener to buttons
      document
        .getElementById("richtigButton")
        .addEventListener("click", karteRichtig);
      document
        .getElementById("falschButton")
        .addEventListener("click", karteFalsch);
    }
  };
  xhr.send(null);
}
function karteSkippen() {
  location.reload();
}
// Functions for button karte richtig/falsch
function karteRichtig() {
  let karteid = this.parentNode.id;
  // Get karte
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/karteikarte/${karteid}`;
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let fach = xhr.response.fach;
      if (fach < 4) {
        fach++;
      }
      // Update current karte
      let xhr2 = new XMLHttpRequest();
      let url2 = `http://localhost:8080/karteikarte/update`;
      xhr2.open("POST", url2, true);
      xhr2.onreadystatechange = function() {
        if (xhr2.readyState == 4 && xhr2.status == 200) {
          location.href = `lern.html?Use_Id=${id}`;
        }
      };

      let data = JSON.stringify({
        karteid: karteid,
        titel: xhr.response.titel,
        frage: xhr.response.titel,
        antwort: xhr.response.antwort,
        fach: fach.toString()
      });
      xhr2.send(data);
    }
  };
  xhr.send(null);
}
function karteFalsch() {
  let karteid = this.parentNode.id;
  // Get karte
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/karteikarte/${karteid}`;
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let fach = xhr.response.fach;
      if (fach < 4) {
        fach++;
      }
      // Update current karte
      let xhr2 = new XMLHttpRequest();
      let url2 = `http://localhost:8080/karteikarte/update`;
      xhr2.open("POST", url2, true);
      xhr2.onreadystatechange = function() {
        if (xhr2.readyState == 4 && xhr2.status == 200) {
          location.href = `lern.html?Use_Id=${id}`;
        }
      };
      let data = JSON.stringify({
        karteid: karteid,
        titel: xhr.response.titel,
        frage: xhr.response.titel,
        antwort: xhr.response.antwort,
        fach: "0"
      });
      xhr2.send(data);
    }
  };
  xhr.send(null);
}

// LOGOUT
document.getElementById("logoutButton").addEventListener("click", logout);

function logout() {
  let xhr = new XMLHttpRequest();
  let url = "http://localhost:8080/login/logout";
  xhr.open("GET", url, true);
  // Callback function
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      window.location.href = "http://localhost:8080/static/";
    }
  };
  // GET data
  xhr.send(null);
}

// Function to decide which fach will be shown to the user
function zufallsFach() {
  let r = Math.floor(Math.random() * (14 - 0 + 1) + 0);
  let f;

  switch (r) {
    case 0:
      f = 4;
      break;
    case 1:
    case 2:
      f = 3;
      break;
    case 3:
    case 4:
    case 5:
      f = 2;
      break;
    case 6:
    case 7:
    case 8:
    case 9:
      f = 1;
      break;
    case 10:
    case 11:
    case 12:
    case 13:
    case 14:
      f = 0;
      break;
  }
  return f;
}

// Add eventListener to new kartei button
document.getElementById("newKarteiButton").addEventListener("click", newKartei);
// NEW KARTEI
function newKartei() {
  location.href = "edit.html";
}
