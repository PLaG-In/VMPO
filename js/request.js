var SERVER_URL = "http://localhost:8080";
var secret = "";
var userID = 0;
var currentTask = [];

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
				secret = result.SecretCode;
				userID = result.User_id;
				alert('Залогинился');
			},
			error: function(){
				pageSetup.showLoginControls();
				console.log("[LOGIN] Unhandled server error");
			}
		});
	return true;
}

function waitWhileEmpty(){
	return currentTask;
}

function addTask(){
	var url = SERVER_URL + '/add_task';
	var date = getDate();
	var task = createTask();
	var priority = "0";
	if (task[3] == "Низкий"){
		priority = "0";
	} else {
		priority = "1";
	}
	var postData = { Secret : secret, user_id : userID, name : task[1], date : date, time : task[2], priority : priority, description : "0" };
	$.ajax({
		type: 'POST',
		url: url,
		data: postData,
		dataType: 'json',
		success: function(result){
			var tr = '<tr>';
			task.forEach(function(item) {
				tr += '<td>' + item + '</td>';
			});
			tr += '</tr>';
			$('#taskTable > tbody:last-child').append(tr);
			getList();
		}
	});
}

function deleteTaskReq(){
	var url = SERVER_URL + '/remove_task';
	var date = getDate();
	var id_task = getTaskId();
	var postData = { Secret : secret, id_task: id_task, id_user : userID, date : date };
		$.ajax({
		type: 'POST',
		url: url,
		data: postData,
		dataType: 'json',
		success: function(result){
			alert("Удалено");
		}
	});
}

function editTask(){
	var url = SERVER_URL + '/edit_task';
	var date = getDate();
	var task = createTask();
	var priority = "0";
	if (task[3] == "Низкий"){
		priority = "0";
	} else {
		priority = "1";
	}
	var postData = { Secret : secret, user_id : userID, name : task[1], date : date, time : task[2], priority : priority, description : " ", task_id: task[0] };
	$.ajax({
		type: 'POST',
		url: url,
		data: postData,
		dataType: 'json',
		success: function(result){
			var tr = '<tr>';
			task.forEach(function(item) {
				tr += '<td>' + item + '</td>';
			});
			tr += '</tr>';
			$('#taskTable > tbody:last-child').append(tr);
		}
	});
}

function getTaskList(){
	var url = SERVER_URL + '/get_list_data';
	var task = [];
	var date = getDate();
	var postData = { secret : secret, date : date, id_user : userID };
	$.ajax({
		type: 'GET',
		url: url,
		data: postData,
		dataType: 'json',
		success: function(result){
			currentTask = result.Task;
		}
	});
}

function exitRequest(){
	
	var url = SERVER_URL + '/sign_out';
	var postData = { login : $('#userName').val(), command : 'sign_out'};
	console.log(postData);
	pageSetup.showLoginControls();
	window.location.reload();
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
				secret = result.SecretCode;
				userID = result.User_id;
				alert('Зарегистрировался');
			},
			error: function(){
				pageSetup.showLoginControls();
				console.log("[SIGNUP] Unhandled server error");
			}
		});
	return true;
}
