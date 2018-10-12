import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Config } from '../../api/config'
import { App } from './app';

@Component({
  selector: 'app-about',
  templateUrl: './about.page.html',
  styleUrls: ['./about.page.scss'],
})
export class AboutPage {
  app: App = {}
  constructor(
    private http: HttpClient,
  ) {
    this.http.get<App>(`${Config.URL}/about`)
      .subscribe(app => this.app = app)
  }
}
