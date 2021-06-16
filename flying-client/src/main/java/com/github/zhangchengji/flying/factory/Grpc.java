package com.github.zhangchengji.flying.factory;
import com.github.zhangchengji.flying.client.Client;
import com.github.zhangchengji.flying.common.Response;
import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.POST;

import java.util.Observable;

public interface Grpc {

    @POST("/client/config")
    Call<Response> config(@Body Client client);
   // boolean  listener();
}
