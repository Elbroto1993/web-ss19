"use strict";

document
  .getElementById("karteikastenInputSelect")
  .addEventListener("change", selectListener);

function selectListener() {
  let kategorie = this.value;
  window.location.href = `http://localhost:8080/meinekarteien?_kategorie=${kategorie}`;
}

let deleteButtons = document.getElementsByClassName("kastenDeleteButton");
for (let i = 0; i < deleteButtons.length; i++) {
  deleteButtons[i].addEventListener("click", deleteKasten);
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
    .addEventListener("click", deleteKasten);
  function deleteKasten() {
    // Delete profile
    let xhr = new XMLHttpRequest();
    let url = "http://localhost:8080/delete-kasten";
    xhr.open("DELETE", url, true);
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200) {
        window.location.href = "http://localhost:8080/meinekarteien";
      }
    };
    xhr.send(JSON.stringify(id));
  }
}
