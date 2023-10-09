<!-- vim: set ft=html: -->
{{template "base" .}}

{{define "content"}}
<div class="container">
    <div class="row">
        <div class="col">
            <h1 class="mt-3">Home page</h1>
            <form action="login" method="post">
                <div class="form-group">
                    <label for="email">Email address</label>
                    <input type="email" class="form-control" id="email" aria-describedby="emailHelp" placeholder="Enter email" name="email">
                </div>
                <div class="form-group">
                    <label for="password">Password</label>
                    <input type="password" class="form-control" id="password" placeholder="Password" name="password">
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>

            <hr/>
            <small>Your request came from {{.IP}}</small>
        </div>
    </div>
</div>
{{end}}
