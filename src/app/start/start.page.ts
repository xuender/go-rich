import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

import { Config } from '../api/config'

@Component({
  selector: 'app-start',
  templateUrl: './start.page.html',
  styleUrls: ['./start.page.scss'],
})
export class StartPage implements OnInit {

  obj = { nick: '', pass: '' }
  constructor(
    private router: Router,
    private http: HttpClient,
  ) {
  }

  ngOnInit() {
  }

  login() {
    this.http.get<string>(`${Config.URL}/login?nick=${this.obj.nick}&pass=${this.obj.pass}`)
      .subscribe(t => {
        localStorage.setItem('token', t)
        this.router.navigateByUrl('/tabs')
      })
  }
}
