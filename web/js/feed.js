const form = $(".leave-comment");

$(document).ready(function() {

    $(".like_button").click(function(e) {
        e.preventDefault();

       console.log(this.closest(".post").id)

        $.ajax({
            type: "POST",
            url: "/like",
            data: JSON.stringify({
                "post_id" : parseInt(this.closest(".post").id)
            }),
            dataType: "json",
            contentType : "application/json",
            success: () => {
                console.log("fail")
                const likes = this.parentNode.querySelector(".like_count")

                if (this.classList.contains("dislike")) {
                    likes.innerHTML = (parseInt(likes.innerHTML) + 1).toString()
                } else {
                    likes.innerHTML = (parseInt(likes.innerHTML) - 1).toString()
                }

                this.classList.toggle("dislike")
                this.classList.toggle("like")
            },
            error: () => {

            }
        });
    });

    $('form').each(function() {
        $(this).validate({

            errorPlacement: function(error, element) {
                error.appendTo(element.closest('form'))
            },

            rules: {
                comment: {
                    required: true,
                    maxlength: 200,
                },
            },

            submitHandler: function(data){
                console.log($(data.querySelector(".comment-input")).val())

                $.ajax({
                    type: "POST",
                    url: "/comment",
                    data: JSON.stringify({
                        "post_id" : parseInt(data.closest(".post").id),
                        "comment" : $(data.querySelector(".comment-input")).val(),
                    }),
                    dataType: "json",
                    contentType : "application/json",
                    success: () => {

                    },
                    error: () => {

                    },
                });
            }
        })
    });
});