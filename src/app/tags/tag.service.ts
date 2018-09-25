import { ModalController } from '@ionic/angular';
import { Injectable, NgZone } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { pull, find } from 'lodash'

import { Tag } from './tag';
import { Key } from '../api/key';
import { URL } from '../api/init'
import { SelectPage } from '../tags/select/select.page';

@Injectable({
  providedIn: 'root'
})
export class TagService {
  private tags: Tag[] = []
  keys: Key[] = []
  constructor(
    private modalCtrl: ModalController,
    private http: HttpClient,
    private ngZone: NgZone,
  ) {
    this.keys.push(
      { key: 'tag-C', name: '客户', color: 'primary', tags: [], },
      { key: 'tag-I', name: '商品', color: 'warning', tags: [], },
    )
  }

  get tags$() {
    return this.http.get<Tag[]>(`${URL}/api/tags`)
  }

  tagsByKey(key: string, all = true) {
    return this.http.get<Tag[]>(`${URL}/api/tags/${key}?all=${all}`)
  }

  search(txt: string) {
    const params = {}
    if (txt) {
      params['search'] = txt
    }
    return this.http.get<Tag[]>(`${URL}/api/tags`, { params: params })
  }

  has(tags: Tag[]): boolean {
    if (tags) {
      this.tags = tags
      return tags.length > 0
    }
    return false;
  }

  async save(tag: Tag) {
    if (tag.id) {
      const t = await this.http.put<Tag>(`${URL}/api/tags/${tag.id}`, tag).toPromise()
      Object.assign(find(this.tags, (o) => o.id == t.id), t)
      return Object.assign(tag, t)
    } else {
      const t = await this.http.post<Tag>(`${URL}/api/tags`, tag).toPromise()
      this.tags.push(t)
      return t
    }
  }

  async delete(tag: Tag) {
    await this.http.delete(`${URL}/api/tags/${tag.id}`).toPromise()
    pull(this.tags, tag)
    this.ngZone.run(() => false)
  }

  async select(key: string) {
    const k = find(this.keys, k => k.key == key)
    const modal = await this.modalCtrl.create({
      component: SelectPage,
      componentProps: {
        tags: k.tags,
        tags$: this.tagsByKey(key),
      }
    });
    await modal.present()
    const d = await modal.onDidDismiss()
    if (d.data) {
      k.tags = d.data
      return d.data
    }
    return null
  }
}
