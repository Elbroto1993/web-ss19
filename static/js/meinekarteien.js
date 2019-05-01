"use strict";

// Global variable to keep track of site state, to reset site after search
let oldDivValue;

document.addEventListener("DOMContentLoaded", displayAllData);

function displayAllData() {
  // Get all kaesten and show them either for own kaesten or public kaesten
  let xhr = new XMLHttpRequest();
  let url = "http://localhost:8080/karteikasten/user";
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let ownKasten = document.getElementById("ownKasten");
      let otherKasten = document.getElementById("kastenFromOtherUser");
      let listKasten = xhr.response;
      let divToAdd;
      for (let i = 0; i < listKasten.length; i++) {
        // Get karten count for every kasten
        let kartenCount;
        let xhr2 = new XMLHttpRequest();
        let url2 = `http://localhost:8080/karteikarte/kasten/${
          listKasten[i].kastenid
        }`;
        xhr2.responseType = `json`;
        xhr2.open("GET", url2, true);
        xhr2.onreadystatechange = function() {
          if (xhr2.readyState == 4 && xhr2.status == 200) {
            if (xhr2.response) {
              kartenCount = xhr2.response.length;
            } else {
              kartenCount = 0;
            }
            // Get string "Sichtbarkeit" for inside kasten to show
            let sichtbarkeit;
            if (listKasten[i].private == "true") {
              sichtbarkeit = "Privat";
            } else {
              sichtbarkeit = "Öffentlich";
            }
            if (listKasten[i].userid === listKasten[i].createdByUserId) {
              divToAdd = `
          <div class="karteikastenCard" id="${listKasten[i].kastenid}">
          <h3 class="karteikastenCardUnterkategorie">
            ${listKasten[i].ueberkategorie}>${listKasten[i].kategorie}
          </h3>
          <div class="karteikastenFlex">
            <a
              href="view.html?Use_Id=${listKasten[i].kastenid}"
              class="karteikastenCardHeader"
              ><strong>${listKasten[i].titel}</strong></a
            >
            <div>
              <div class="hexagon hexagon2">
                <div class="hexagon-in1">
                  <div class="hexagon-in2"></div>
                </div>
              </div>
              <div class="karteikastenInsideHexagon">
                <p id="karteikastenCounterKartenHexagon">
                  ${kartenCount}
                </p>
                <p>Karten</p>
              </div>
            </div>
          </div>
          <p class="karteikastenCardBeschreibung">
            ${listKasten[i].beschreibung}
          </p>
          <p class="karteikastenCardBeschreibung">
            Sichtbarkeit: <strong id="sichtbarkeit">${sichtbarkeit}</strong>
          </p>
          <p class="karteikastenCardBeschreibung">
            Fortschritt: <strong id="fortschritt">${
              listKasten[i].fortschritt
            }%</strong>
          </p>
          <button class="btnYellow kastenLernButton">Lernen</button>
          <button class="btnGreen kastenEditButton">Bearbeiten</button>
          <button class="btnRed kastenDeleteButton"><img src="icons/Delete.svg"></button>
        </div>
        `;
              ownKasten.innerHTML += divToAdd;
            } else {
              divToAdd = `
        <div class="karteikastenCard" id="${listKasten[i].kastenid}">
          <h3 class="karteikastenCardUnterkategorie">${
            listKasten[i].ueberkategorie
          }>${listKasten[i].kategorie}</h3>
          <div class="karteikastenFlex">
            <a
              href="view.html?Use_Id=${listKasten[i].kastenid}"
              class="karteikastenCardHeader"
              ><strong>${listKasten[i].titel}</strong></a
            >
            <div>
              <div class="hexagon hexagon2">
                <div class="hexagon-in1">
                  <div class="hexagon-in2"></div>
                </div>
              </div>
              <div class="karteikastenInsideHexagon">
                <p id="karteikastenCounterKartenHexagon">
                  ${kartenCount}
                </p>
                <p>Karten</p>
              </div>
            </div>
          </div>
          <button class="btnYellow kastenLernButtonPublic">Lernen</button>
        </div>
        `;
              otherKasten.innerHTML += divToAdd;
            }
            // Fill oldDivValue to reset site when search bar is empty
            oldDivValue = document.getElementById("kaesten").innerHTML;
            ////////////////
          }

          // Add event listener for all buttons
          let lernButtons = document.getElementsByClassName("kastenLernButton");
          for (let i = 0; i < lernButtons.length; i++) {
            lernButtons[i].addEventListener("click", lernKasten);
          }
          let editButtons = document.getElementsByClassName("kastenEditButton");
          for (let i = 0; i < editButtons.length; i++) {
            editButtons[i].addEventListener("click", editKasten);
          }
          let deleteButtons = document.getElementsByClassName(
            "kastenDeleteButton"
          );
          for (let i = 0; i < deleteButtons.length; i++) {
            deleteButtons[i].addEventListener("click", deleteKasten);
          }
          let lernButtonsPublic = document.getElementsByClassName(
            "kastenLernButtonPublic"
          );
          for (let i = 0; i < lernButtonsPublic.length; i++) {
            lernButtonsPublic[i].addEventListener("click", lernKasten);
          }
        };
        xhr2.send(null);
      }
    }
  };
  xhr.send(null);
  // Get kasten counter
  // Prepare http request
  let xhr3 = new XMLHttpRequest();
  xhr3.responseType = `json`;
  url = "http://localhost:8080/karteikasten/";
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

// Functions for all kasten buttons
function lernKasten() {
  let id = this.parentNode.id;
  location.href = `lern.html?Use_Id=${id}`;
}
function editKasten() {
  let id = this.parentNode.id;
  location.href = `edit2.html?Use_Id=${id}`;
}
function deleteKasten() {
  let id = this.parentNode.id;
  // Toggle modal for second delete button
  document.getElementById("profilDeleteModal").classList.add("is-active");

  document
    .getElementById("modal-close")
    .addEventListener("click", untoggleDeleteModal);
  document
    .getElementById("profilKeepButton")
    .addEventListener("click", untoggleDeleteModal);

  function untoggleDeleteModal() {
    document.getElementById("profilDeleteModal").classList.remove("is-active");
  }
  // Final delete button
  document
    .getElementById("profilFinallyDeleteButton")
    .addEventListener("click", deleteProfile);
  function deleteProfile() {
    // Delete kasten
    let xhr = new XMLHttpRequest();
    let url = `http://localhost:8080/karteikasten/${id}`;
    xhr.open("DELETE", url, true);
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200) {
        location.reload();
      }
    };
    xhr.send(null);
  }
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
        if (karteikaesten[i].kategorie == kategorie) {
          divChild.innerHTML += `
          <div class="karteikastenCard">
            <h3 class="karteikastenCardUnterkategorie">${
              karteikaesten[i].kategorie
            }</h3>
            <a href="#" class="karteikastenCardHeader"
              ><strong>${karteikaesten[i].titel}</strong></a
            >
            <p class="karteikastenCardBeschreibung">
            ${karteikaesten[i].beschreibung}
            </p>
            <button class="btnYellow">Lernen</button>
          </div>`;
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
      divChild.innerHTML += `
      <div class="karteikastenCard">
        <h3 class="karteikastenCardUnterkategorie">${
          karteikaesten[i].kategorie
        }</h3>
        <a href="#" class="karteikastenCardHeader"
          ><strong>${karteikaesten[i].titel}</strong></a
        >
        <p class="karteikastenCardBeschreibung">
        ${karteikaesten[i].beschreibung}
        </p>
        <button class="btnYellow">Lernen</button>
      </div>`;
    }
  }
  if (!inputValue) {
    divToAdd.innerHTML = oldDivValue;
  }
}

// Add eventListener to new kartei button
document.getElementById("newKarteiButton").addEventListener("click", newKartei);
// NEW KARTEI
function newKartei() {
  location.href = "edit.html";
}
