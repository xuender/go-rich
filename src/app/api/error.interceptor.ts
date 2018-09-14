import { Injectable } from "@angular/core";
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from "@angular/common/http";
import { Observable } from "rxjs";
import { tap } from "rxjs/operators";
import { ProfileService } from "./profile.service";

@Injectable()
export class ErrorInterceptor implements HttpInterceptor {
  constructor(
    private profile: ProfileService,
  ) { }
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(req).pipe(tap((event: HttpEvent<any>) => { }, (err: any) => {
      switch (err.status) {
        case 400:
        case 401:
          this.profile.login()
          break
        default:
          this.profile.error(err.error.message)
      }
    }))
  }
}
