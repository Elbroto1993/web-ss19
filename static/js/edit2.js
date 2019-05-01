"use strict";

// MDE EDITOR
let simplemde = new SimpleMDE({
  element: document.getElementById("edit2Textarea1")
});

let simplemde2 = new SimpleMDE({
  element: document.getElementById("edit2Textarea2")
});

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
    }
  };
  xhr.send(null);
  // Get all karten pro kasten
  let xhr2 = new XMLHttpRequest();
  let url3 = `http://localhost:8080/karteikarte/kasten/${id}`;
  xhr2.responseType = `json`;
  xhr2.open("GET", url3, true);
  xhr2.onreadystatechange = function() {
    if (xhr2.readyState == 4 && xhr2.status == 200) {
      let data = xhr2.response;
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
          <div class="edit2Grid">
            <p class="edit2Number">#${i + 1}</p>
            <h4 class="edit2Titel"><strong>${data[i].titel}</strong></h4>
            <div class="edit2Buttons" id="${data[i].karteid}">
              <button class="btnGreen editButton">Bearbeiten</button>
              <button class="btnRed edit2deleteButton"><img src="icons/Delete.svg"></button>
            </div>
          </div>
        </div>
       `;
          kartListe.innerHTML += divToAdd;
        }
        // Add event listener for every karte in the list
        let editButtons = document.getElementsByClassName("editButton");
        for (let i = 0; i < editButtons.length; i++) {
          editButtons[i].addEventListener("click", editKarte);
        }
        let edit2deleteButtons = document.getElementsByClassName(
          "edit2deleteButton"
        );
        for (let i = 0; i < edit2deleteButtons.length; i++) {
          edit2deleteButtons[i].addEventListener("click", deleteKarte);
        }
      }
    }
  };
  xhr2.send(null);
}
// Functions for buttons when no karte is selected
document.getElementById("saveButton").addEventListener("click", saveKarte);
function saveKarte() {
  // Get values from single card forms
  let titel = document.getElementById("edit2input").value;
  let frage = simplemde.value();
  let antwort = simplemde2.value();
  // Prepare http post request
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/karteikarte/`;
  xhr.open("POST", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      console.log(xhr.responseText);
    }
  };
  let data = JSON.stringify({
    kastenid: id,
    titel: titel,
    frage: frage,
    antwort: antwort,
    fach: "0"
  });
  xhr.send(data);
}
// Function for edit kasten button
document
  .getElementById("editKastenButton")
  .addEventListener("click", editKasten);

function editKasten() {
  location.href = `edit.html?Use_Id=${id}`;
}

document.getElementById("stopButton").addEventListener("click", stopEdit);
function stopEdit() {
  window.location.href = "http://localhost:8080/static/meinekarteien.html";
}
// Functions for each button in kartenlist
function editKarte() {
  let id = this.parentNode.id;
  // Fill elements for single card
  let singleCard = document.getElementById("edit2SingleCard");
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/karteikarte/${id}`;
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let karte = xhr.response;
      let divToAdd = `
      <div class="edit2SingleCardGrid">
      <div class="viewPositionTitle">Titel</div>
      <input type="text" class="viewPositionTitle" value="${
        karte.titel
      }" id="edit2input"/>
      <div>Frage</div>
      <div>
        <textarea
          name=""
          id="edit2Textarea1"
          cols="100"
          rows="10"
        >${karte.frage}</textarea>
      </div>
      <div>Antwort</div>
      <div>
        <textarea
          name=""
          id="edit2Textarea2"
          cols="100"
          rows="7"
        >${karte.antwort}</textarea>
      </div>
      </div>
      <div class="edit2SingleCardButtons" id="${id}">
        <button class="btnYellow updateButton">
         Speichern
        </button>
        <button class="btnRed abortButton">
          Abbrechen
        </button>
      </div>
        `;
      singleCard.innerHTML = divToAdd;
      // Add event listener to buttons from single card
      let updateButtons = document.getElementsByClassName("updateButton");
      for (let i = 0; i < updateButtons.length; i++) {
        updateButtons[i].addEventListener("click", updateCard);
      }
      let abortButtons = document.getElementsByClassName("abortButton");
      for (let i = 0; i < abortButtons.length; i++) {
        abortButtons[i].addEventListener("click", abortCard);
      }
    }
  };
  xhr.send(null);
}
// Functions for single card buttons
function updateCard(e) {
  e.preventDefault();
  // Get values from form
  let titel = document.getElementById("edit2input").value;
  let frage = document.getElementById("edit2Textarea1").value;
  let antwort = document.getElementById("edit2Textarea2").value;
  // Prepare http post request
  let id = this.parentNode.id;
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/karteikarte/update`;
  xhr.open("POST", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      location.reload();
    }
  };
  let data = JSON.stringify({
    karteid: id,
    titel: titel,
    frage: frage,
    antwort: antwort
  });
  xhr.send(data);
}
// Show empty input fields for single card
function abortCard() {
  document.getElementById("edit2SingleCard").innerHTML = `
  <div class="edit2SingleCardGrid">
  <div class="viewPositionTitle">Titel</div>
  <input type="text" class="viewPositionTitle" id="edit2input"/>
  <div>Frage</div>
  <div>
    <textarea
      name=""
      id="edit2Textarea1"
      cols="100"
      rows="10"
    ></textarea>
  </div>
  <div>Antwort</div>
  <div>
    <textarea
      name=""
      id="edit2Textarea2"
      cols="100"
      rows="7"
    ></textarea>
  </div>
  </div>
  <div class="edit2SingleCardButtons">
    <button class="btnYellow" id="saveButton">
     Speichern
   </button>
   <button class="btnRed" id="stopButton">
      Abbrechen
   </button>
  </div>
  `;
  document.getElementById("saveButton").addEventListener("click", saveKarte);
  function saveKarte() {
    // Get values from single card forms
    let titel = document.getElementById("edit2input").value;
    let frage = document.getElementById("edit2Textarea1").value;
    let antwort = document.getElementById("edit2Textarea2").value;
    // Prepare http post request
    let xhr = new XMLHttpRequest();
    let url = `http://localhost:8080/karteikarte/`;
    xhr.open("POST", url, true);
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200) {
        location.reload();
      }
    };
    let data = JSON.stringify({
      kastenid: id,
      titel: titel,
      frage: frage,
      antwort: antwort,
      fach: "0"
    });
    xhr.send(data);
  }
  document.getElementById("stopButton").addEventListener("click", stopEdit);
  function stopEdit() {
    window.location.href = "http://localhost:8080/static/meinekarteien.html";
  }
}
function deleteKarte() {
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
    .getElementById("profilFinallyedit2deleteButton")
    .addEventListener("click", deleteKarte);
  function deleteKarte() {
    // Delete kasten
    let xhr = new XMLHttpRequest();
    let url = `http://localhost:8080/karteikarte/${id}`;
    xhr.open("DELETE", url, true);
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200) {
        location.reload();
      }
    };
    xhr.send(null);
  }
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

// Add eventListener to new kartei button
document.getElementById("newKarteiButton").addEventListener("click", newKartei);
// NEW KARTEI
function newKartei() {
  location.href = "edit.html";
}
