"use strict";

document
  .getElementById("karteikastenInputSelect")
  .addEventListener("change", selectListener);

function selectListener() {
  let kategorie = this.value;
  window.location.href = `http://localhost:8080/karteikasten?_kategorie=${kategorie}`;
  // let xhr = new XMLHttpRequest();
  // let url = `http://localhost:8080/karteikasten`;
  // xhr.open("POST", url, true);
  // xhr.onreadystatechange = function() {
  //   if (xhr.readyState == 4 && xhr.status == 200) {
  //   }
  // };
  // let data = JSON.stringify({
  //   kategorie: kategorie
  // });
  // xhr.send(data);
  //   let ueberkategorie = this.options[this.selectedIndex].parentNode.getAttribute(
  //     "label"
  //   );
  //   let naturwissenschaften = document.getElementById("naturwissenschaften");
  //   let sprachen = document.getElementById("sprachen");
  //   let wirtschaft = document.getElementById("wirtschaft");
  //   let geisteswissenschaften = document.getElementById("geisteswissenschaften");
  //   sprachen.style.display = "block";
  //   wirtschaft.style.display = "block";
  //   geisteswissenschaften.style.display = "block";
  //   naturwissenschaften.style.display = "block";
  //   switch (ueberkategorie) {
  //     case "Naturwissenschaften":
  //       sprachen.style.display = "none";
  //       wirtschaft.style.display = "none";
  //       geisteswissenschaften.style.display = "none";
  //       break;
  //     case "Sprachen":
  //       naturwissenschaften.style.display = "none";
  //       wirtschaft.style.display = "none";
  //       geisteswissenschaften.style.display = "none";
  //       break;
  //     case "Wirtschaft":
  //       sprachen.style.display = "none";
  //       naturwissenschaften.style.display = "none";
  //       geisteswissenschaften.style.display = "none";
  //       break;
  //     case "Geisteswissenschaften":
  //       sprachen.style.display = "none";
  //       wirtschaft.style.display = "none";
  //       naturwissenschaften.style.display = "none";
  //       break;
  //     default:
  //       break;
  //   }
}
