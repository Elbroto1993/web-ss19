"use strict";

// Get kastenid from url
let urlParam = function(name, w) {
  w = w || window;
  var rx = new RegExp("[&|?]" + name + "=([^&#]+)"),
    val = w.location.search.match(rx);
  return !val ? "" : val[1];
};
let kastenid = urlParam("_kastenid");

// STORE NEW KASTEN
document
  .getElementById("editKastenButton")
  .addEventListener("click", newKasten);

function newKasten(e) {
  e.preventDefault();
  // Get data from forms
  let kategorieError = document.getElementById("kategoriePflicht");
  let titel = document.getElementById("editKastenTitel").value;
  let kategorieHelp = document.getElementById("editKastenSelect");
  let kategorie = kategorieHelp.options[kategorieHelp.selectedIndex].value;
  let beschreibung = document.getElementById("editKastenBeschreibung").value;
  let ueberkategorie = kategorieHelp.options[
    kategorieHelp.selectedIndex
  ].parentNode.getAttribute("label");
  // Clear errors
  kategorieError.innerHTML = "";
  kategorieError.classList.remove("registerAlert");
  // Get value from radio buttons
  let privateHelp = document.getElementsByClassName("editRadioButton");
  let privateOrPublic;
  for (let i = 0; i < privateHelp.length; i++) {
    if (privateHelp[i].checked) {
      privateOrPublic = privateHelp[i].value;
    }
  }
  // Change value to boolean
  if (privateOrPublic === "Öffentlich") {
    privateOrPublic = "false";
  } else {
    privateOrPublic = "true";
  }
  // Check if kategorie is selected
  if (!kategorie) {
    kategorieError.innerHTML = "Pflichtfeld!";
    kategorieError.classList.add("registerAlert");
  } else {
    let xhr = new XMLHttpRequest();
    let url = "http://localhost:8080/add-or-update-kasten";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200) {
        let data = xhr.response.replace(/"/g, "");
        window.location.href = `http://localhost:8080/edit2?_kastenid=${data}`;
      }
    };
    // Set data for POST
    let data = JSON.stringify({
      _id: kastenid,
      titel: titel,
      kategorie: kategorie,
      beschreibung: beschreibung,
      private: privateOrPublic,
      ueberkategorie: ueberkategorie
    });
    xhr.send(data);
    // Clear forms
    document.getElementById("editKastenTitel").value = "";
    document.getElementById("editKastenBeschreibung").value = "";
  }
}
