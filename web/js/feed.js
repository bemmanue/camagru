const form = $(".leave-comment");

$(document).ready(function() {

    $(".like_button").click(function(e) {
        e.preventDefault();

        $.ajax({
            type: "POST",
            url: "/like",
            data: JSON.stringify({
                "post_id" : parseInt(this.closest(".post").id)
            }),
            dataType: "json",
            contentType : "application/json",
            success: () => {
                const likes = this.parentNode.querySelector(".like_count")

                if (this.classList.contains("dislike")) {
                    likes.innerHTML = (parseInt(likes.innerHTML) + 1).toString()
                } else {
                    likes.innerHTML = (parseInt(likes.innerHTML) - 1).toString()
                }

                this.classList.toggle("dislike")
                this.classList.toggle("like")
            },
            error: () => {}
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
                let comment_text = $(data.querySelector(".comment-input")).val()
                data.querySelector(".comment-input").value = ""

                $.ajax({
                    type: "POST",
                    url: "/comment",
                    data: JSON.stringify({
                        "post_id" : parseInt(data.closest(".post").id),
                        "comment" : comment_text,
                    }),
                    dataType: "json",
                    contentType : "application/json",
                    success: () => {
                        let ul = data.parentNode.parentNode.querySelector('.comments')

                        let commentCount = data.parentNode.parentNode.querySelector(".comment_count")
                        commentCount.innerHTML = (parseInt(commentCount.innerHTML) + 1).toString()

                        let b = document.createElement("b");
                        b.appendChild(document.createTextNode(localStorage.getItem("username")));

                        let li = document.createElement("li");
                        li.setAttribute("class", "comment")
                        li.appendChild(b);
                        li.appendChild(document.createTextNode('\u00A0'));
                        li.appendChild(document.createTextNode(comment_text));

                        ul.appendChild(li)
                    },
                    error: () => {},
                });
            }
        })
    });
});