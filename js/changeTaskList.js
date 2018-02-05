var selectedTask = null;
var currentTaskId = 0;
var taskList = {
	1: [1, "Умри", "00:00:00", "Высокий"],
	2: [2, "Убей", "00:00:00", "Низкий"],
	3: [3, "Живи", "00:00:00", "Высокий"],
	4: [4, "Кричи", "00:00:00", "Высокий"],
};

function addTask()
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
	currentTaskId = selectedTask.context.cells[0].textContent;
}

function getTaskTime()
{
	return selectedTask.context.cells[2].textContent;
}

function putTime()
{
	if (isSelected())
	{
		$('#taskTable tbody tr')[currentTaskId - 1].cells[2].textContent = getTime();
		if (taskList[currentTaskId] !== undefined)
		{
			taskList[currentTaskId][2] = getTime();
		}else{
			alert("Вы удалили задачу, которую выполняли(не надо так)")
		}
	}
}

function deselectTask(){
	selectedTask = null;
}

function deleteTask()
{
	if (selectedTask != null)
	{
		var id = selectedTask.context.cells[0].textContent;
		delete taskList[id];
		selectedTask.remove();
		deselectTask();
	}
}

function changeTask()
{
	var id = selectedTask.context.cells[0].textContent;
	taskList[id][1] = document.getElementById('eNameOfTask').value;
	selectedTask.context.cells[1].textContent = taskList[id][1];
	taskList[id][3] = document.getElementById('ePriority').value;
	selectedTask.context.cells[3].textContent = taskList[id][3];
}

function getTask()
{
	if (selectedTask != null)
	{
		document.getElementById('eNameOfTask').value = selectedTask.context.cells[1].textContent;
		//description get
		document.getElementById('ePriority').value = selectedTask.context.cells[3].textContent;
	}
}

function getTaskLista()
{
	return taskList;
}

function createTask() {
	var id = 0;
	var task = $('#aNameOfTask').val();
	var time = "00:00:00";
	var priority = $('#aPriority').val();
	return [id, task, time, priority];
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
