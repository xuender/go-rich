import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from './user';
import { URL } from './init'
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  users: Observable<User[]>
  constructor(
    private http: HttpClient
  ) {
    this.users = this.http.get<User[]>(`${URL}/api/users`)
  }
}
