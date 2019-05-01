"use strict";

// Global variables to reduce http requests
let userKaesten;
let userId;

// Get kastenid from url
let urlParam = function(name, w) {
  w = w || window;
  var rx = new RegExp("[&|?]" + name + "=([^&#]+)"),
    val = w.location.search.match(rx);
  return !val ? "" : val[1];
};
let id = urlParam("Use_Id");

document.addEventListener("DOMContentLoaded", displayAllData);

function displayAllData() {
  // CHECK IF USER IS LOGGED IN
  let xhr5 = new XMLHttpRequest();
  let url2 = "http://localhost:8080/login/checkLoggedIn";
  xhr5.responseType = `json`;
  xhr5.open("GET", url2, true);
  xhr5.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr5.onreadystatechange = function() {
    if (xhr5.readyState == 4 && xhr5.status == 200) {
      // If user is logged in add sidebar items and change login field
      if (xhr5.response) {
        // Add sidebar items
        document.getElementById("menuList").innerHTML += `<li class="menuItems">
          <a href="meinekarteien.html"
            ><img
              src="icons/Meine-Karteien.svg"
              alt=""
              height="30"
              width="30"
            />
            Meine Karteien<small class="karteiZaehler" id="indexKartenCounterUser">0</small></a
          >
          </li>
          <li class="menuItems">
            <a href="profil.html"
              ><img
                src="icons/Mein-Profil.svg"
                alt=""
                height="30"
                width="30"
              />
              Mein Profil</a
            >
          </li>`;
        // Change login field
        document.getElementById("loginField").innerHTML = `<div class="logout">
          <button class="btnYellow newKarteiButton" id="newKarteiButton">
            Neue Kartei
          </button>
          <div class="headerUserInfo">
            <h4 id="logoutUsername">Max Mustermann</h4>
            <button class="logoutButton" id="logoutButton">
             <small>Logout</small>
            </button>
          </div>
            <img
            src="images/defaultUser.png"
            alt="gray default male profile picture"
            height="30"
            width="30"
            />
          </div>`;
        // Change name in logout field
        // Prepare http request
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
            // Add userid value to global variable
            userId = res.id;
          }
        };
        // GET data
        xhr6.send(null);
        // Add eventListener to logout button
        document
          .getElementById("logoutButton")
          .addEventListener("click", logout);
        // Add eventListener to new kartei button
        document
          .getElementById("newKarteiButton")
          .addEventListener("click", newKartei);
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
              // Add kaesten value to global variable
              userKaesten = xhr4.response;
            }
          }
        };
        // GET data
        xhr4.send(null);
      }
    }
  };
  xhr5.send(null);
  // This also happens for users who aren't logged in
  // Get kasten counter for sidebar
  let indexKastenCounterAll = document.getElementById("indexKastenCounterAll");
  // Prepare http request
  let xhr3 = new XMLHttpRequest();
  xhr3.responseType = `json`;
  let url3 = "http://localhost:8080/karteikasten/";
  xhr3.open("GET", url3, true);
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
  // Get all infos about kasten
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/karteikasten/kasten/${id}`;
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      // Get current user to check if kasten is from user
      let xhr7 = new XMLHttpRequest();
      let url4 = `http://localhost:8080/user/id`;
      xhr7.responseType = `json`;
      xhr7.open("GET", url4, true);
      xhr7.onreadystatechange = function() {
        if (xhr7.readyState == 4 && xhr7.status == 200) {
          let data = xhr.response;
          // Fill elements (name, kategorie, fortschritt) with data from kasten
          document.getElementById("viewName").innerHTML = data.titel;
          document.getElementById("viewKategorie").innerHTML =
            data.ueberkategorie + ">" + data.kategorie;
          // Only add fortschritt if current user owns the kasten
          if (xhr7.response) {
            if (xhr7.response.id == data.userid) {
              document.getElementById("viewFortschritt").innerHTML =
                data.fortschritt + " %";
            } else {
              document.getElementById("viewFortschritt").innerHTML = "0 %";
            }
          }
          // Get username for createdById
          let xhr2 = new XMLHttpRequest();
          let url2 = `http://localhost:8080/user/id/${data.createdByUserId}`;
          xhr2.responseType = `json`;
          xhr2.open("GET", url2, true);
          xhr2.onreadystatechange = function() {
            if (xhr2.readyState == 4 && xhr2.status == 200) {
              let data = xhr2.response;
              // Fill element (erstellt von) with data from createdByUserId
              document.getElementById("viewCreatedBy").innerHTML =
                data.username;
            }
          };
          xhr2.send(null);
        }
      };
      xhr7.send(null);
    }
  };
  xhr.send(null);
}

// Get all karten pro kasten
let xhr3 = new XMLHttpRequest();
let url3 = `http://localhost:8080/karteikarte/kasten/${id}`;
xhr3.responseType = `json`;
xhr3.open("GET", url3, true);
xhr3.onreadystatechange = function() {
  if (xhr3.readyState == 4 && xhr3.status == 200) {
    let data = xhr3.response;
    // Fill element (anzahl karten) with data from length of slice, if data isn't null
    if (!data) {
      document.getElementById("viewCountKarten").innerHTML = 0;
    } else {
      document.getElementById("viewCountKarten").innerHTML = data.length;

      // List all karten
      let kartListe = document.getElementById("viewListKarten");
      for (let i = 0; i < data.length; i++) {
        let divToAdd = `
        <div class="card">
          <p>#${i + 1}</p>
          <button class="noStyleButton karteButton" id="${data[i].karteid}">${
          data[i].titel
        }</button>
        </div>
       `;
        kartListe.innerHTML += divToAdd;
      }
      // Add event listener for every karte in the list
      let karten = document.getElementsByClassName("karteButton");
      for (let i = 0; i < karten.length; i++) {
        karten[i].addEventListener("click", displayKarte);
      }
    }
  }
};
xhr3.send(null);

function displayKarte() {
  // Remove green background from old selected card
  let karteHelp = document.getElementsByClassName("card");
  for (let i = 0; i < karteHelp.length; i++) {
    karteHelp[i].classList.remove("viewBackgroundGreen");
  }
  // Remove white font color from buttons
  let buttonHelp = document.getElementsByClassName("karteButton");
  for (let i = 0; i < buttonHelp.length; i++) {
    buttonHelp[i].classList.remove("viewKarteButtonWhiteFont");
  }
  // Add green background for new selected card
  this.parentNode.classList += " viewBackgroundGreen";
  this.classList += " viewKarteButtonWhiteFont";
  // Fill elements for single card
  let singleCard = document.getElementById("viewSingleCard");
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/karteikarte/${this.id}`;
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let karte = xhr.response;
      let divToAdd = `
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
      `;

      singleCard.innerHTML = divToAdd;
      let karteFromCurrentUser = false;
      for (let i = 0; i < userKaesten.length; i++) {
        if (userKaesten[i].kastenid == karte.kastenid) {
          karteFromCurrentUser = true;
        }
      }
      if (karteFromCurrentUser) {
        // Add green background to hexagon fach if karte is owned by current user
        let hexToAddOne = document
          .getElementById("singleCardDiv")
          .getElementsByClassName("viewHexagon");
        let hexToAddTwo = document
          .getElementById("singleCardDiv")
          .getElementsByClassName("hexagon-in2");
        for (let i = 0; i < 5; i++) {
          if (karte.fach == i) {
            hexToAddOne[i].id = "hex1";
            hexToAddTwo[i].id = "hex2";
          }
        }
      }
    }
  };
  xhr.send(null);
}

// LOGIN
document.getElementById("loginButton").addEventListener("click", login);

function login(e) {
  e.preventDefault();
  // Get values from form
  let username = document.getElementById("loginUsername").value;
  let password = document.getElementById("loginPassword").value;
  // Check if both inputs are filled
  if (username === "" || password === "") {
    alert("Bitte beide Felder ausfüllen.");
  } else {
    // Prepare http request
    let xhr4 = new XMLHttpRequest();
    let url = "http://localhost:8080/login/";
    xhr4.open("POST", url, true);
    xhr4.responseType = `json`;
    xhr4.setRequestHeader("Content-Type", "application/json");
    // Callback function
    xhr4.onreadystatechange = function() {
      if (xhr4.readyState == 4 && xhr4.status == 200) {
        if (xhr4.response) {
          window.location.href =
            "http://localhost:8080/static/meinekarteien.html";
        } else {
          alert("Nutzername oder Passwort sind falsch!");
        }
      }
    };
    // POST data
    let data = JSON.stringify({
      username: username,
      password: password
    });
    xhr4.send(data);
    // Clear forms after submit
    document.getElementById("loginUsername").value = "";
    document.getElementById("loginPassword").value = "";
  }
}

// LOGOUT
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

// NEW KARTEI
function newKartei() {
  location.href = "edit.html";
}
