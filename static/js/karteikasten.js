"use strict";

document
  .getElementById("karteikastenInputSelect")
  .addEventListener("change", selectListener);

function selectListener() {
  let kategorie = this.value;
  window.location.href = `http://localhost:8080/karteikasten?_kategorie=${kategorie}`;
}
