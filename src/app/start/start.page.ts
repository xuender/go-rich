import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { Config } from '../api/config'
import { App } from '../setting/about/app';

@Component({
  selector: 'app-start',
  templateUrl: './start.page.html',
  styleUrls: ['./start.page.scss'],
})
export class StartPage implements OnInit {

  obj = { nick: '', pass: '' }
  app: App
  constructor(
    private router: Router,
    private http: HttpClient,
  ) {
    this.http.get<App>(`${Config.URL}/about`)
      .subscribe(app => this.app = app)
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
