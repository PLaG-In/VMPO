		// Get the modal
var signup = document.getElementById('id01');
var login = document.getElementById('id02');
var add = document.getElementById('id03');
var edit = document.getElementById('id04');

	// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
	if (event.target == signup) {
		signup.style.display = "none";
	} else if (event.target == login){
		login.style.display = "none";
	} else if(event.target == add){
		add.style.display = "none";
	} else	if(event.target == edit){
		edit.style.display = "none"
	}

}