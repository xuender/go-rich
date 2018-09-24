import { Injectable, NgZone } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { User } from './user';
import { ObjService } from '../api/obj.service';

@Injectable({
  providedIn: 'root'
})
export class UserService extends ObjService<User>{
  constructor(
    protected http: HttpClient,
    protected ngZone: NgZone,
  ) {
    super(http, ngZone)
  }
  path() { return '/api/users' }
}
