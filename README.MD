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



# Resources
    -   handeling html form data in AWS API Gateway (https://forums.aws.amazon.com/thread.jspa?messageID=673012&tstart=0#673012)
    -   go lambda logging (https://medium.com/@tjholowaychuk/apex-log-e8d9627f4a9a#.jjwfjcv90)