import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from './user';
import { URL } from './init'

@Injectable({
  providedIn: 'root'
})
export class UserService {

  users: User[] = []
  constructor(
    private http: HttpClient
  ) {
    this.load()
  }
  load() {
    this.http.get<User[]>(`${URL}/api/users`)
      .subscribe((us: User[]) => {
        this.users = us
      })
  }
}
