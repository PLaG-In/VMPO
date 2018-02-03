var selectedTask = null;
var taskList = {
	1: [1, "������", "12.12.2012", "18:00"],
	2: [2, "����� ������", "13.13.2013", "13:00"],
	3: [3, "����� ���� ���������", "10.02.2018", "04:20"],
	4: [4, "������ � ����������� � ������ � ������", "11.02.2018", "00:00"],
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
	var tr = '<tr>'; // ������� ������ �������
	task.forEach(function(item) {
		tr += '<td>' + item + '</td>'; // ��������� ������� � ������
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
	//��� ����� ������
	var id = 0;
	var task = "�����������";
	var date = "30.02.2018";
	var time = "20:00";
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