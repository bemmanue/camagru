$(document).ready(function() {

    jQuery.validator.addMethod("login", function(value, element) {
        return this.optional(element) || /^[a-zA-Z\d]+$/.test(value);
    }, "Only alphanumeric characters");

    jQuery.validator.addMethod("password", function(value, element) {
        return this.optional(element) || /^[ -~]+$/.test(value);
    }, "Only printable characters");


    $("form").validate({
        rules: {
            login: {
                required: true,
                login: true,
                minlength: 6,
                maxlength: 30,
            },
            email: {
                required: true,
                email: true
            },
            password: {
                required: true,
                password: true,
                minlength: 6,
                maxlength: 30
            },
            password_confirm: {
                required: true,
                equalTo: "#password"
            }
        }
    })
});
