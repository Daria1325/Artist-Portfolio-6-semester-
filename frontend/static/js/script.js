// var expanded = false;
//
// function showCheckboxes() {
//     var checkboxes = document.getElementById("checkboxes");
//     if (!expanded) {
//         checkboxes.style.display = "block";
//         expanded = true;
//     } else {
//         checkboxes.style.display = "none";
//         expanded = false;
//     }
// }

var anyFormOpen = false;

function openForm(s, objID, objName, objDesc, objPath, objPrice, objDate, objMat, objSize) {
  objType = s.split('_')[1];


  if (!anyFormOpen) {
        document.getElementById(s).style.visibility = "visible";
        anyFormOpen = true;
        if(objType=='s'){
            if(arguments.length == 1){//add
              document.querySelector('#' + s).querySelector('form').action = 'series/add';
            }
            else if (arguments.length == 2) {//del
              document.querySelector('#' + s).querySelector('a').href += objID;
            } else if (arguments.length > 3) {//edit
              document.querySelector('#' + s).querySelector('form').action = 'series/edit/' + objID;
              // fill Form
              document.querySelector('#' + s).querySelector('#id').value = objID ;
              document.querySelector('#' + s).querySelector('#name').value = objName.trim() ;
              document.querySelector('#' + s).querySelector('#description').value = objDesc ;

            }
        }
        else {
              if(arguments.length == 1){//add
                document.querySelector('#' + s).querySelector('form').action = 'pictures/add';
              }else if (arguments.length == 2) {//del
                  document.querySelector('#' + s).querySelector('a').href += objID;
              }else if (arguments.length > 3) {//edit
                  document.querySelector('#' + s).querySelector('form').action = 'pictures/edit/' + objID;
console.log(document.querySelector('#' + s).querySelector('form').action)
                  document.querySelector('#' + s).querySelector('#id').value = objID ;
                  document.querySelector('#' + s).querySelector('#name').value = objName.trim() ;
                  document.querySelector('#' + s).querySelector('#description').value = objDesc.trim() ;
                  document.querySelector('#' + s).querySelector('#size').value = objSize.trim() ;
                  document.querySelector('#' + s).querySelector('#price').value = objPrice ;
                  document.querySelector('#' + s).querySelector('#year').value = objDate ;
                  document.querySelector('#' + s).querySelector('#material').value = objMat.trim() ;
                  // document.querySelector('#' + s).querySelector('#img').value = objPath.trim();
                  // document.querySelector('#' + s).querySelector('#series').value = objPath.trim();


              }
        }
  }

}

function closeForm(s) {
  objType = s.split('_')[1];
  document.getElementById(s).style.visibility = "hidden";
  anyFormOpen = false;

  document.querySelector('#' + s).querySelector('form').method = 'get';
  if (objType == s){
    document.querySelector('#' + s).querySelector('form').action = 'series';
  }else {
    document.querySelector('#' + s).querySelector('form').action = 'pictures';
  }





}



