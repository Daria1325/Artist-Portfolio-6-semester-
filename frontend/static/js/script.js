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

function openForm(s, objID, objName, objDesc) {
  if (!anyFormOpen) {
    document.getElementById(s).style.visibility = "visible";
    anyFormOpen = true;
    if(arguments.length == 1){//add
      document.querySelector('#' + s).querySelector('form').action = 'series/add';
      // console.log(document.querySelector('#' + s).querySelector('form').action);
    }
    else if (arguments.length == 2) {//del
      document.querySelector('#' + s).querySelector('a').href += objID;
      console.log(document.querySelector('#' + s).querySelector('a').href);

    } else if (arguments.length == 4) {//edit
      document.querySelector('#' + s).querySelector('form').action = 'series/edit/' + objID;
      // console.log(document.querySelector('#' + s).querySelector('form').action);

      // fill Form
      document.querySelector('#' + s).querySelector('#id').value = objID ;
      document.querySelector('#' + s).querySelector('#name').value = objName.trim() ;
      document.querySelector('#' + s).querySelector('#description').value = objDesc ;

      console.log(document.querySelector('#' + s).querySelector('#name'));

    }
  }

}

function closeForm(s) {
  document.getElementById(s).style.visibility = "hidden";
  anyFormOpen = false;

  document.querySelector('#' + s).querySelector('form').method = 'get';
  document.querySelector('#' + s).querySelector('form').action = 'series';




}

// TODO
// value="{{(index .Items $key).ID }}"
// action="series/edit/objID"

// close add


