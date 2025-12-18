import { describe, it, expect } from 'vitest'
import { cn } from '../utils'

describe('utils', () => {
  describe('cn', () => {
    it('合并类名', () => {
      expect(cn('foo', 'bar')).toBe('foo bar')
    })

    it('处理条件类名', () => {
      expect(cn('foo', false && 'bar', 'baz')).toBe('foo baz')
    })

    it('处理undefined和null', () => {
      expect(cn('foo', undefined, null, 'bar')).toBe('foo bar')
    })
  })
})
