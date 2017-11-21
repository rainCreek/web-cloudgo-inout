$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(data) {
       $('.greeting-student').append(data.student);
       $('.greeting-studentid').append(data.studentid);
       $('.greeting-title').append(data.title);
    });
});


