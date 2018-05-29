    var currentMonth = "0";
    var currentYear = 2018;
	var currentDate = "0";
	function Calendar2(id, year, month) {
    	var Dlast = new Date(year,month+1,0).getDate(),
    	    D = new Date(year,month,Dlast),
    	    DNlast = new Date(D.getFullYear(),D.getMonth(),Dlast).getDay(),
    	    DNfirst = new Date(D.getFullYear(),D.getMonth(),1).getDay(),
    	    calendar = '<tr>',
    	    month=["Январь","Февраль","Март","Апрель","Май","Июнь","Июль","Август","Сентрябрь","Октябрь","Ноябрь","Декабрь"];
    	if (DNfirst != 0)
    	{
    		for(var  i = 1; i < DNfirst; i++) calendar += '<td>';
    	}
    	else
    	{
      		for(var  i = 0; i < 6; i++) calendar += '<td>';
    	}

    	for(var  i = 1; i <= Dlast; i++) {
      		if (i == new Date().getDate() && D.getFullYear() == new Date().getFullYear() && D.getMonth() == new Date().getMonth())
      		{
        		calendar += '<td class="cell today">' + i;
      		}
      		else
      		{
        		calendar += '<td class = "cell">' + i;
      		}
      		if (new Date(D.getFullYear(),D.getMonth(),i).getDay() == 0)
      		{
        		calendar += '<tr>';
      		}
    	}
    	for(var  i = DNlast; i < 7; i++) calendar += '<td>&nbsp;';
    	document.querySelector('#'+id+' tbody').innerHTML = calendar;
    	document.querySelector('#'+id+' thead td:nth-child(2)').innerHTML = month[D.getMonth()] +' '+ D.getFullYear();
    	document.querySelector('#'+id+' thead td:nth-child(2)').dataset.month = D.getMonth();
    	document.querySelector('#'+id+' thead td:nth-child(2)').dataset.year = D.getFullYear();
    	if (document.querySelectorAll('#'+id+' tbody tr').length < 6)
    	{  // ����� ��� �������������� ������� �� "������������" ��� ��������, ����������� ��� ������ ������. ����: ������ 6 ����� ��� ����
        	document.querySelector('#'+id+' tbody').innerHTML += '<tr><td>&nbsp;<td>&nbsp;<td>&nbsp;<td>&nbsp;<td>&nbsp;<td>&nbsp;<td>&nbsp;';
    	}
    }

    Calendar2("calendar2", new Date().getFullYear(), new Date().getMonth());
    // ������������� ����� �����
    document.querySelector('#calendar2 thead tr:nth-child(1) td:nth-child(1)').onclick = function() {
      Calendar2("calendar2", document.querySelector('#calendar2 thead td:nth-child(2)').dataset.year, parseFloat(document.querySelector('#calendar2 thead td:nth-child(2)').dataset.month)-1);
    }
    // ������������� ���� �����
    document.querySelector('#calendar2 thead tr:nth-child(1) td:nth-child(3)').onclick = function() {
      Calendar2("calendar2", document.querySelector('#calendar2 thead td:nth-child(2)').dataset.year, parseFloat(document.querySelector('#calendar2 thead td:nth-child(2)').dataset.month)+1);
    }

	function getDate(){
		if (Number(currentDate) < 10 && currentDate[0] != 0){
			currentDate = "0" + currentDate;
		}
		if (Number(currentMonth) < 10 && currentMonth[0] != 0){
			currentMonth = "0" + currentMonth
		}
		var date = currentMonth + "." + currentDate + "." + currentYear;
		return date;
	}

	function getList(){	
		
		//setTimeout(2000);
		deselectTask();
		getTaskList();
		setTimeout(function(){
			var taskList = waitWhileEmpty();
			callBack(taskList);
		}, 200);
		 
		
	 
		//getTaskLists(newTaskList);
	
	}

	function callBack(taskList){
		var newTaskList = [];
		getTaskLists(newTaskList);
		var task = [];
		// if (taskList === undefined) {
		// 	taskList = [];
		// };
		$("#taskTable tbody tr").detach();
    	for (i = 0; i < taskList.length; i++) {
		  var tr = '<tr>';
		  var z = 0;
		  for (item in taskList[i])
		  { 
			task[z] = taskList[i][item];
			z++;
			tr += '<td>' + taskList[i][item] + '</td>';
		  }
		  newTaskList[i] = task;
		  tr += '</tr>';
		  $('#taskTable > tbody:last-child').append(tr);
		}
	}
	
	$(document).ready(function()
	{
	    $('table').on('click', '.cell', function ()
	    {
			var index = -1;
			if ($('.click').index() != -1){
				index = $('.click')[0].innerText;
			}	
			$('.cell').filter(function(){
				return $(this)[0].innerText == index;
			}).removeClass('click');
			currentDate = $(this)[0].innerText;
			currentMonth = Number($('table thead tr td')[1].dataset.month) + 1;
			currentYear = $('table thead tr td')[1].dataset.year;
			$("#taskTable tbody tr").detach();
			getList();
			
			
			//alert($('.click').index());
			$(this).toggleClass('click');
	    });
	});
