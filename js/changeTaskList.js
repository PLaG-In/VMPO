var selectedTask = null;
var currentTaskId = 0;
var taskList = {

};

function addTasks()
{
	var task = createTask();
	var lastKey = 0;
	for (var key in taskList) {
		lastKey = key;
	}
	var newKey = Number(lastKey) + 1;
	task[0] = newKey;
	taskList[newKey] = task;
	var tr = '<tr>';
	task.forEach(function(item) {
		tr += '<td>' + item + '</td>';
	});
	tr += '</tr>';
	$('#taskTable > tbody:last-child').append(tr);
}

function isSelected() {
	if (selectedTask != null)
	{
		return true;
	}
	return false;
}

function chooseTask()
{
	currentTaskId = selectedTask.index();
}

function getTaskTime()
{
	return selectedTask.context.cells[2].textContent;
}

function putTime()
{
	if (isSelected())
	{
		$('#taskTable tbody tr')[currentTaskId].cells[2].textContent = getTime();
		if (taskList[currentTaskId] !== undefined)
		{
			taskList[currentTaskId][2] = getTime();
			updateTime();
		}else{
			alert("Вы удалили задачу, которую выполняли(не надо так)")
		}
	}
}

function deselectTask(){
	selectedTask = null;
}

function getTaskId(){
	return selectedTask.context.cells[0].textContent;
}

function deleteTask()
{
	if (selectedTask != null)
	{
		deleteTaskReq();
		var id = getTaskId();
		delete taskList[id];
		selectedTask.remove();
		deselectTask();
	}
}

function changeTask()
{
	var id = selectedTask.context.cells[0].textContent;
	//taskList[id][1] = document.getElementById('eNameOfTask').value;
	selectedTask.context.cells[1].textContent = document.getElementById('eNameOfTask').value;
	//taskList[id][3] = document.getElementById('ePriority').value;
	selectedTask.context.cells[3].textContent = document.getElementById('eDescriptionOfTask').value;
	editTask(selectedTask.context.cells[1].textContent, selectedTask.context.cells[3].textContent);
}

function getTask()
{
	if (selectedTask != null)
	{
		document.getElementById('eNameOfTask').value = selectedTask.context.cells[1].textContent;
		//description get
		document.getElementById('eDescriptionOfTask').value = selectedTask.context.cells[3].textContent;
	}
}

function getTaskLists(newTaskList){
	taskList = newTaskList;
}

function updateTask() {
	var task_id = getTaskId();
	var task = document.getElementById('eNameOfTask').value = selectedTask.context.cells[1].textContent;
	var time = getTime();
	var des = document.getElementById('eDescriptionOfTask').value = selectedTask.context.cells[3].textContent;
	return [task_id, task, time, des];
}

function createTask() {
	var id = 0;
	var task = $('#aNameOfTask').val();
	var time = "00:00:00";
	var des = $('#aDescriptionOfTask').val();
	return [id, task, time, des];
}

$(document).ready(function()
	{
	    $('#taskTable').on('click', 'tbody tr', function ()
	    {
				if (selectedTask != $(this) && selectedTask != null)
				{
					selectedTask.removeClass('selectlines');
				}
				selectedTask = $(this);
				$(this).toggleClass('selectlines');
				getTime();
	    });
	});
