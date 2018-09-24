import { Component, OnInit } from '@angular/core';

import { SettingService } from './setting.service';
import { ExtService } from '../ext/ext.service';

@Component({
  selector: 'app-setting',
  templateUrl: './setting.page.html',
  styleUrls: ['./setting.page.scss'],
})
export class SettingPage implements OnInit {

  constructor(
    public setting: SettingService,
    public ext: ExtService
  ) { }

  ngOnInit() {
  }

}
