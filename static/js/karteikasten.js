"use strict";

// Global variable to keep track of site state, to reset site after search
let oldDivValue;
let loggedIn;

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
      loggedIn = xhr5.response;
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
        // Add eventListener to logout button
        document
          .getElementById("logoutButton")
          .addEventListener("click", logout);
        // Add eventListener to new kartei button
        document
          .getElementById("newKarteiButton")
          .addEventListener("click", newKartei);
        let xhr6 = new XMLHttpRequest();
        let url = "http://localhost:8080/user/id";
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
            if (xhr4.response) {
              indexKartenCounterUser.innerHTML = xhr4.response.length;
            }
          }
        };
        // GET data
        xhr4.send(null);
      }
    }
  };
  xhr5.send(null);
  let karteikaesten;
  // THIS ALSO HAPPENS FOR USERS WHO AREN'T LOGGED IN
  // Prepare http request
  let xhr3 = new XMLHttpRequest();
  xhr3.responseType = `json`;
  let url = "http://localhost:8080/karteikasten/";
  xhr3.open("GET", url, true);
  xhr3.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr3.onreadystatechange = function() {
    if (xhr3.readyState == 4 && xhr3.status == 200) {
      // Save public kaesten in global variable
      karteikaesten = xhr3.response;
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
      // DISPLAY ALL PUBLIC KAESTEN
      let naturwissenschaften = document.getElementById("naturwissenschaften");
      let sprachen = document.getElementById("sprachen");
      let wirtschaft = document.getElementById("wirtschaft");
      let geisteswissenschaften = document.getElementById(
        "geisteswissenschaften"
      );
      let divToAdd;
      let numberOfCards;
      for (let i = 0; i < karteikaesten.length; i++) {
        if (karteikaesten[i].private === "false") {
          // Get number of karten inside kasten
          let xhr7 = new XMLHttpRequest();
          xhr7.responseType = `json`;
          url = `http://localhost:8080/karteikarte/kasten/${
            karteikaesten[i].kastenid
          }`;
          xhr7.open("GET", url, true);
          // Callback function
          xhr7.onreadystatechange = function() {
            if (xhr7.readyState == 4 && xhr7.status == 200) {
              if (xhr7.response != undefined) {
                numberOfCards = xhr7.response.length;
              } else {
                numberOfCards = 0;
              }
              // If logged in, add lernen button else don't
              if (loggedIn) {
                // String for every div
                divToAdd = `          
                    <div class="karteikastenCard">
                      <h3 class="karteikastenCardUnterkategorie">${
                        karteikaesten[i].kategorie
                      }</h3>              
                      <div class="karteikastenFlex">
                        <a href="view.html?Use_Id=${
                          karteikaesten[i].kastenid
                        }" class="karteikastenCardHeader"
                          ><strong>${karteikaesten[i].titel}</strong></a
                        >
                       <div>
                         <div class="hexagon hexagon2">
                          <div class="hexagon-in1">
                            <div class="hexagon-in2"></div>
                          </div>
                        </div>
                       <div class="karteikastenInsideHexagon">
                         <p id="karteikastenCounterKartenHexagon">${numberOfCards}</p>
                          <p>Karten</p>
                        </div>
                      </div>
                    </div>
                    <p class="karteikastenCardBeschreibung">
                    ${karteikaesten[i].beschreibung}
                    </p>
                    <button class="btnYellow lernenButtons" id="${
                      karteikaesten[i].kastenid
                    }">Lernen</button>
                    </div>`;
              } else {
                divToAdd = `          
                    <div class="karteikastenCard">
                      <h3 class="karteikastenCardUnterkategorie">${
                        karteikaesten[i].kategorie
                      }</h3>              
                      <div class="karteikastenFlex">
                        <a href="view.html?Use_Id=${
                          karteikaesten[i].kastenid
                        }" class="karteikastenCardHeader"
                          ><strong>${karteikaesten[i].titel}</strong></a
                        >
                      <div>
                        <div class="hexagon hexagon2">
                        <div class="hexagon-in1">
                          <div class="hexagon-in2"></div>
                        </div>
                      </div>
                      <div class="karteikastenInsideHexagon">
                        <p id="karteikastenCounterKartenHexagon">${numberOfCards}</p>
                        <p>Karten</p>
                      </div>
                      </div>
                    </div>
                    <p class="karteikastenCardBeschreibung">
                    ${karteikaesten[i].beschreibung}
                    </p>
                  </div>`;
              }
              // Switch to decide ueberkategorie
              switch (karteikaesten[i].ueberkategorie) {
                case "Naturwissenschaften":
                  naturwissenschaften.innerHTML += divToAdd;
                  break;
                case "Sprachen":
                  sprachen.innerHTML += divToAdd;
                  break;
                case "Wirtschaft":
                  wirtschaft.innerHTML += divToAdd;
                  break;
                case "Geisteswissenschaften":
                  geisteswissenschaften.innerHTML += divToAdd;
                  break;
              }
              // Fill oldDivValue to reset site when search bar is empty
              oldDivValue = document.getElementById("kaesten").innerHTML;
              ////////////////
              let lernenButtons = document.getElementsByClassName(
                "lernenButtons"
              );
              for (let j = 0; j < lernenButtons.length; j++) {
                if (lernenButtons[j]) {
                  lernenButtons[j].addEventListener("click", lernenAddKasten);
                }
              }
            }
          };
          xhr7.send(null);
        }
      }
    }
  };
  // GET data
  xhr3.send(null);
}

// LERNEN BUTTON ADD KASTEN TO USER
function lernenAddKasten() {
  let thisKastenId = this.id;
  let redirectKastenId;
  // Get kasten from lernen button
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/karteikasten/kasten/${thisKastenId}`;
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      // Check if user already is learning the selected kasten
      // Get all kaesten from logged in user
      let xhr5 = new XMLHttpRequest();
      let url5 = `http://localhost:8080/karteikasten/user`;
      xhr5.responseType = `json`;
      xhr5.open("GET", url5, true);
      xhr5.onreadystatechange = function() {
        if (xhr5.readyState == 4 && xhr5.status == 200) {
          let isLearningKasten = false;
          let res = xhr5.response;
          // Check if user is already using this kasten
          if (res) {
            for (let i = 0; i < res.length; i++) {
              if (
                xhr.response.kategorie == res[i].kategorie &&
                xhr.response.titel == res[i].titel &&
                xhr.response.beschreibung == res[i].beschreibung
              ) {
                isLearningKasten = true;
                redirectKastenId = res[i].kastenid;
              }
            }
          }
          // If he uses it, redirect to lern.html
          if (isLearningKasten) {
            location.href = `lern.html?Use_Id=${redirectKastenId}`;
            // Else copy kasten to user and then redirect
          } else {
            // Store kasten for user by cookie user id
            let xhr2 = new XMLHttpRequest();
            let url2 = `http://localhost:8080/karteikasten/`;
            xhr2.responseType = `json`;
            xhr2.open("POST", url2, true);
            xhr2.onreadystatechange = function() {
              if (xhr2.readyState == 4 && xhr2.status == 200) {
                let newKastenId = xhr2.response.kastenid;
                redirectKastenId = newKastenId;
                // Get all karten from old kasten
                let xhr8 = new XMLHttpRequest();
                let url5 = `http://localhost:8080/karteikarte/kasten/${thisKastenId}`;
                xhr8.responseType = `json`;
                xhr8.open("GET", url5, true);
                xhr8.onreadystatechange = function() {
                  if (xhr8.readyState == 4 && xhr8.status == 200) {
                    let response = xhr8.response;
                    // Loop through every karte and store it to new kasten
                    for (let i = 0; i < response.length; i++) {
                      let xhr4 = new XMLHttpRequest();
                      let url4 = `http://localhost:8080/karteikarte/`;
                      xhr4.responseType = `json`;
                      xhr4.open("POST", url4, true);
                      xhr4.onreadystatechange = function() {
                        if (xhr4.readyState == 4 && xhr4.status == 200) {
                          console.log(xhr4.response);
                        }
                      };
                      let data = JSON.stringify({
                        kastenid: newKastenId,
                        titel: response[i].titel,
                        frage: response[i].frage,
                        antwort: response[i].antwort,
                        fach: "0"
                      });
                      xhr4.send(data);
                      location.href = `lern.html?Use_Id=${redirectKastenId}`;
                    }
                  }
                };
                xhr8.send(null);
              }
            };
            // Create json struct for post request
            let data = JSON.stringify({
              createdbyuserid: xhr.response.createdByUserId,
              kategorie: xhr.response.kategorie,
              fortschritt: xhr.response.forstschritt,
              private: "true",
              beschreibung: xhr.response.beschreibung,
              fortschritt: "0",
              ueberkategorie: xhr.response.ueberkategorie,
              titel: xhr.response.titel
            });
            xhr2.send(data);
          }
        }
      };
      xhr5.send(null);
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

// SORTIEREN NACH...
document
  .getElementById("karteikastenInputSelect")
  .addEventListener("change", sortierenNach);

function sortierenNach() {
  let xhr = new XMLHttpRequest();
  let url = "http://localhost:8080/karteikasten/";
  xhr.open("GET", url, true);
  xhr.responseType = `json`;
  xhr.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let karteikaesten = xhr.response;
      let kategorieHelp = document.getElementById("karteikastenInputSelect");
      let kategorie = kategorieHelp.options[kategorieHelp.selectedIndex].value;
      let ueberkategorie = kategorieHelp.options[
        kategorieHelp.selectedIndex
      ].parentNode.getAttribute("label");

      let divParent = document.getElementById("kaesten");
      divParent.innerHTML = `
      <div>
        <h2 class="karteikastenThemaHeader">
          <strong>${ueberkategorie}</strong>
        </h2>
      </div>
      <div class="flex karteikastenLayout" id="${ueberkategorie}"></div>`;
      let divChild = document.getElementById(`${ueberkategorie}`);
      for (let i = 0; i < karteikaesten.length; i++) {
        if (
          karteikaesten[i].kategorie == kategorie &&
          karteikaesten[i].private == "false"
        ) {
          if (loggedIn) {
            divChild.innerHTML += `          
            <div class="karteikastenCard">
              <h3 class="karteikastenCardUnterkategorie">${
                karteikaesten[i].kategorie
              }</h3>
              <a href="view.html?Use_Id=${
                karteikaesten[i].kastenid
              }" class="karteikastenCardHeader"
                ><strong>${karteikaesten[i].titel}</strong></a
              >
              <p class="karteikastenCardBeschreibung">
              ${karteikaesten[i].beschreibung}
              </p>
              <button class="btnYellow lernenButtons" id="${
                karteikaesten[i].kastenid
              }">Lernen</button>
            </div>`;
          } else {
            divChild.innerHTML += `          
            <div class="karteikastenCard">
              <h3 class="karteikastenCardUnterkategorie">${
                karteikaesten[i].kategorie
              }</h3>
              <a href="view.html?Use_Id=${
                karteikaesten[i].kastenid
              }" class="karteikastenCardHeader"
                ><strong>${karteikaesten[i].titel}</strong></a
              >
              <p class="karteikastenCardBeschreibung">
              ${karteikaesten[i].beschreibung}
              </p>
            </div>`;
          }
        }
      }
      let lernenButtons = document.getElementsByClassName("lernenButtons");
      if (lernenButtons) {
        for (let j = 0; j < lernenButtons.length; j++) {
          if (lernenButtons[j]) {
            lernenButtons[j].addEventListener("click", lernenAddKasten);
          }
        }
      }
    }
  };
  xhr.send(null);
}

// SUCHE
let karteikaesten;
let xhr = new XMLHttpRequest();
let url = "http://localhost:8080/karteikasten/";
xhr.open("GET", url, true);
xhr.responseType = `json`;
xhr.setRequestHeader("Content-Type", "application/json");
// Callback function
xhr.onreadystatechange = function() {
  if (xhr.readyState == 4 && xhr.status == 200) {
    karteikaesten = xhr.response;
  }
};
xhr.send(null);
// Add event listener to search field that listens to every letter input
document
  .getElementById("karteikastenInputSearch")
  .addEventListener("keyup", searchKaesten);

function searchKaesten() {
  let inputValue = document.getElementById("karteikastenInputSearch").value;
  let divToAdd = document.getElementById("kaesten");
  divToAdd.innerHTML = `<div class="flex karteikastenLayout" id="divChild"></div>`;
  let divChild = document.getElementById("divChild");
  for (let i = 0; i < karteikaesten.length; i++) {
    if (
      karteikaesten[i].private == "false" &&
      (karteikaesten[i].kategorie.includes(inputValue) ||
        karteikaesten[i].titel.includes(inputValue) ||
        karteikaesten[i].beschreibung.includes(inputValue))
    ) {
      if (loggedIn) {
        divChild.innerHTML += `          
        <div class="karteikastenCard">
          <h3 class="karteikastenCardUnterkategorie">${
            karteikaesten[i].kategorie
          }</h3>
          <a href="view.html?Use_Id=${
            karteikaesten[i].kastenid
          }" class="karteikastenCardHeader"
            ><strong>${karteikaesten[i].titel}</strong></a
          >
          <p class="karteikastenCardBeschreibung">
          ${karteikaesten[i].beschreibung}
          </p>
          <button class="btnYellow lernenButtons" id="${
            karteikaesten[i].kastenid
          }">Lernen</button>
        </div>`;
      } else {
        divChild.innerHTML += `          
        <div class="karteikastenCard">
          <h3 class="karteikastenCardUnterkategorie">${
            karteikaesten[i].kategorie
          }</h3>
          <a href="view.html?Use_Id=${
            karteikaesten[i].kastenid
          }" class="karteikastenCardHeader"
            ><strong>${karteikaesten[i].titel}</strong></a
          >
          <p class="karteikastenCardBeschreibung">
          ${karteikaesten[i].beschreibung}
          </p>
        </div>`;
      }
    }
  }
  let lernenButtons = document.getElementsByClassName("lernenButtons");
  if (lernenButtons) {
    for (let j = 0; j < lernenButtons.length; j++) {
      if (lernenButtons[j]) {
        lernenButtons[j].addEventListener("click", lernenAddKasten);
      }
    }
  }
  if (!inputValue) {
    divToAdd.innerHTML = oldDivValue;
    let lernenButtons = document.getElementsByClassName("lernenButtons");
    if (lernenButtons) {
      for (let j = 0; j < lernenButtons.length; j++) {
        if (lernenButtons[j]) {
          lernenButtons[j].addEventListener("click", lernenAddKasten);
        }
      }
    }
  }
}

// NEW KARTEI
function newKartei() {
  location.href = "edit.html";
}
