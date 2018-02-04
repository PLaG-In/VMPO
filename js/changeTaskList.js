var selectedTask = null;
var taskList = {
	1: [1, "Умри", "18:00", "Высокий"],
	2: [2, "Убей", "13:00", "Низкий"],
	3: [3, "Живи", "04:20", "Высокий"],
	4: [4, "Кричи", "00:00", "Высокий"],
};

function addTask()
{
	var task = getTask();
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

function deleteTask()
{
	if (selectedTask != null)
	{
		var id = selectedTask.context.cells[0].textContent;
		delete taskList[id];
		selectedTask.remove();
	}
}

function changeTask()
{

}

function getTaskList() {
	return taskList;
}

function getTask() {
	var id = 0;
	var task = "Беги";
	var time = "20:00";
	var date = "Высокий";
	return [id, task, date, time];
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
	    });
	});
