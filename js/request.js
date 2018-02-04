var SERVER_URL = "http://localhost:8080";
var secret = "";
var userID = 0;
function loginRequest(){
		var url = SERVER_URL + '/auth';

		var userStr = $('#alogin').val();
		var passwordStr = $('#apassword').val();
		if (userStr == '' || passwordStr == '')
		{
			alert('Please fill all fields');
			return false;
		}
		pageSetup.hideLoginControls();
		$('#userName').text(userStr);

		var postData = { login : userStr, password : passwordStr };
		$.ajax({
			type: 'POST',
			url: url,
			data: postData,
			dataType: 'json',
			success: function(result){
				alert('Залогинился');
			},
			error: function(){
				pageSetup.showLoginControls();
				console.log("[LOGIN] Unhandled server error");
			}
		});
	return true;
}

function getTaskList(){
	var url = SERVER_URL + '/get_list_data';
	// var secret =
	// var date =
	// var id =
	var postData = {/* secret : secret, date : date, id_user : id*/ };
	$.ajax({
		type: 'GET',
		url: url,
		data: postData,
		dataType: 'json',
		success: function(result){
			alert('ты получил список');
		}
	});
	return true;
}

function exitRequest(){
	var url = SERVER_URL + '/sign_out';
	var postData = { login : $('#userName').val(), command : 'sign_out'};
	console.log(postData);
	pageSetup.showLoginControls();
	return $.ajax({
		async: false,
		type: 'POST',
		url: url,
		data: postData,
		dataType: 'json',
		error: function(result){
			console.log("[sign_out] Unhandled server error");
		}
	});
}

function signupRequest(){
		var url = SERVER_URL + '/reg';

		var userStr = $('#slogin').val();
		var passwordStr = $('#spassword').val();
		var repasswordStr = $('#repassword').val();
		if (userStr == '' || passwordStr == '')
		{
			alert('Заполните все поля');
			return false;
		}

		if (passwordStr != repasswordStr)
		{
			alert('Пароли не совпадают');
			return false;
		}
		pageSetup.hideLoginControls();
		$('#userName').text(userStr);
		var postData = { login : userStr, password : passwordStr };
		$.ajax({
			type: 'POST',
			url: url,
			data: postData,
			dataType: 'json',
			success: function(result){
				secret = $(result).filter("SecretCode");
				userID = $(result).filter("User_ID");
				alert('Зарегистрировался');
			},
			error: function(){
				pageSetup.showLoginControls();
				console.log("[SIGNUP] Unhandled server error");
			}
		});
	return true;
}
