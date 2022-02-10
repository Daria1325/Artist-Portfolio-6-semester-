var expanded = false;

function showCheckboxes() {
  var checkboxes = document.getElementById("checkboxes");
  if (!expanded) {
    checkboxes.style.display = "block";
    expanded = true;
  } else {
    checkboxes.style.display = "none";
    expanded = false;
  }
}




var anyFormOpen = false;

function openForm(s,
) {
  if (!anyFormOpen) {
    document.getElementById(s).style.visibility = "visible";
    anyFormOpen = true;
    if (arguments.length == 2) {
      document.querySelector('#' + s).querySelector('a').href += objID;
      console.log(document.querySelector('#' + s).querySelector('a').href);
    }
  }

}

function closeForm(s) {
  document.getElementById(s).style.visibility = "hidden";
  anyFormOpen = false;

}

