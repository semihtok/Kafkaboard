<div class="col-12 grid-margin stretch-card">
    <div class="card">
        <div class="card-body">
            <h4 class="card-title">Topic Name : </h4>
            <div class="form-group">
                <div class="input-group">
                    <div class="input-group-prepend">
                        <span class="input-group-text">@</span>
                    </div>
                    <input id="txtTopic" type="text" class="form-control" placeholder="Write a topic for listening messages" aria-label="Topic"
                        aria-describedby="basic-addon1">
                </div>
            </div>
        </div>
    </div>
</div>


<div class="col-md-12">
    <div class="card">
        <div class="card-body">
            <h5 class="card-title"></h5>
            <div id="realtimeChart" style="min-width: 310px; height: 400px; margin: 0 auto">
                </canvas>
            </div>
        </div>
    </div>
</div>

<div class="col-md-12" style="margin-top:20px !important">
    <div class="card">
        <div class="card-body">
            <h5 class="card-title">Latest Messages</h5>
            <div class="table-responsive table-sales">
                <table class="table latestMessages">
                    <thead>
                        <tr>
                            <th colspan="2">
                                Time
                            </th>

                            <div class="flag">
                                <i class="flag-icon flag-icon-us"></i>
                            </div>
                            <th class="text-right">
                                Context
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td class="text-left"></td>
                            <td>
                                <div class="flag">
                                    <i class="flag-icon flag-icon-tr"></i>
                                </div>
                            </td>

                            <td class="text-right">

                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>