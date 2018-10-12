import { Component, OnInit } from '@angular/core';

import { ExtService } from '../ext/ext.service';
import { SettingService } from './setting.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-setting',
  templateUrl: './setting.page.html',
  styleUrls: ['./setting.page.scss'],
})
export class SettingPage implements OnInit {
  constructor(
    public setting: SettingService,
    public ext: ExtService,
    private router: Router,
  ) { }

  ngOnInit() {
  }

  about() {
    this.router.navigateByUrl('/tabs/(setting:about)')
  }

}
