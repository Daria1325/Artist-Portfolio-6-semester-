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

function openForm(s, objID, key) {
  if (!anyFormOpen) {
    document.getElementById(s).style.visibility = "visible";
    anyFormOpen = true;
    if (arguments.length == 2) {
      document.querySelector('#' + s).querySelector('a').href += objID;
      console.log(document.querySelector('#' + s).querySelector('a').href);
    } else if (arguments.length == 3) {
      document.querySelector('#' + s).querySelector('form').action = 'series/edit/' + objID;
      console.log(document.querySelector('#' + s).querySelector('form').action);

      // fill Form 
      // document.querySelector('#' + s).querySelector('#id').value = '{{(index .Items 6).ID }}';
      console.log(document.querySelector('#' + s).querySelector('#name'));

    }
  }

}

function closeForm(s) {
  document.getElementById(s).style.visibility = "hidden";
  anyFormOpen = false;

}

// TODO
// value="{{(index .Items $key).ID }}"
// action="series/edit/objID"

// close add 
