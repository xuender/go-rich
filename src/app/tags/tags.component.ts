import { Component, OnInit, Input } from '@angular/core';
import { Observable } from 'rxjs';
import { includes, pull } from 'lodash'

import { Tag } from './tag';
import { TagService } from './tag.service';
import { SettingService } from '../setting/setting.service';

@Component({
  selector: 'rich-tags',
  templateUrl: './tags.component.html',
  styleUrls: ['./tags.component.scss']
})
export class TagsComponent implements OnInit {
  @Input()
  key: string = 'tag-C'
  @Input()
  tags: string[] = []
  tags$: Observable<Tag[]>
  constructor(
    private tagService: TagService,
    private settingService: SettingService,
  ) { }

  ngOnInit() {
    this.tags$ = this.tagService.tagsByKey(this.key)
  }

  change(tag:string) {
    if(this.has(tag)){
      pull(this.tags, tag)
    }else {
      this.tags.push(tag)
    }
  }

  has(tag: string) {
    return includes(this.tags, tag)
  }
  async edit(){
    await this.settingService.tags()
    this.tags$ = this.tagService.tagsByKey(this.key)
  }
}
