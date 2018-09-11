import { Injectable } from "@angular/core";
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from "@angular/common/http";
import { Observable } from "rxjs";
import { tap } from "rxjs/operators";
import { AlertController } from "@ionic/angular";

@Injectable()
export class ErrorInterceptor implements HttpInterceptor {
  constructor(
    private alertCtrl: AlertController
  ) { }
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(req).pipe(tap((event: HttpEvent<any>) => { }, (err: any) => {
      if (err.status === 500) {
        this.alertCtrl.create({
          header: '错误',
          message: err.error.message,
          buttons: ['确定']
        }).then(a => a.present())
      }
    }))
  }
}
