$(document).ready(function() {
    //Disable the Remove Button
    var rowCount = $('table >tbody:last >tr').length;
    if(rowCount == 1) {
        document.getElementsByClassName('btn-remove')[0].disabled = true;
    }
    
    $(document).on('click', '.btn-add', function(e) {
        e.preventDefault();

        var controlForm = $('table');
        var currentEntry = $('table>tbody>tr:last');
        currentEntry[0].cells[7].firstElementChild.disabled=false;
        var newEntry = $(currentEntry.clone()).appendTo(controlForm);

        //Add the button  
        var rowCount = $('table >tbody:last >tr').length;
        if(rowCount > 1) {
            var removeButtons = document.getElementsByClassName('btn-remove');
            for(var i = 0; i < removeButtons.length; i++) {
                removeButtons.item(i).disabled = false;
            }
        }
         
    }).on('click', '.btn-remove', function(e) {
        $(this).parents('tr:first').remove();
        
        //Disable the Remove Button
        var rowCount = $('table >tbody:last >tr').length;
        if(rowCount == 1) {
            document.getElementsByClassName('btn-remove')[0].disabled = true;
        }
        e.preventDefault();
        return false;
    });
});

window.
