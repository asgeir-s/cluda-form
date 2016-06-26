    Content-Type: application/x-www-form-urlencoded

    
# if you pay
    -   webhook
    -   hidden email
    -   1000 submissions per month


# usage
        <form action="https://form.cluda.com/submission/your@email.com"method="POST">
            <input type="text" name="name">
            <input type="email" name="email">
            <input type="submit" value="Send">
        </form>

    -   get an confirmation email on first submission
    -   after that submissions are sent to the email (max 25 per month)