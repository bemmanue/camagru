$(document).ready(function() {
    $("form").validate({
        rules: {
            login: {
                required: true,
                minlength: 5,
                maxlength:30
            },
            email: {
                required: true,
                email: true
            },
            password: {
                required: true,
                minlength: 5,
                maxlength: 30
            },
            password_confirm: {
                required: true,
                equalTo: "#password"
            }
        }
    })
});
