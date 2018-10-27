import { Injectable } from "@angular/core";
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from "@angular/common/http";
import { Observable } from "rxjs";
import { tap } from "rxjs/operators";
import { AlertController } from '@ionic/angular';
import { Router } from "@angular/router";

@Injectable()
export class ErrorInterceptor implements HttpInterceptor {
  constructor(
    private alertCtrl: AlertController,
    private router: Router,
  ) { }

  async loginError() {
    const alert = await this.alertCtrl.create({
      header: `登录错误`,
      message: '登录名或密码错误，请重新输入',
      buttons: ['确定']
    })
    await alert.present()
    await alert.onDidDismiss()
    this.router.navigateByUrl('/start')
  }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(req).pipe(tap((event: HttpEvent<any>) => { }, (err: any) => {
      switch (err.status) {
        case 401:
          this.loginError()
          break
        case 403:
          this.router.navigateByUrl('/start')
          break
        default:
          this.alertCtrl.create({
            header: `错误 [ ${err.status} ]`,
            message: err.error.message,
            buttons: ['确定']
          }).then(a => a.present())
      }
    }))
  }
}
