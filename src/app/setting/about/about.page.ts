import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Config } from '../../api/config'

@Component({
  selector: 'app-about',
  templateUrl: './about.page.html',
  styleUrls: ['./about.page.scss'],
})
export class AboutPage {
  app = {}
  constructor(
    private http: HttpClient,
  ) {
    this.http.get<any>(`${Config.URL}/about`)
      .subscribe(app => this.app = app)
  }
}
