const backSpaceCompare = (s: string, t: string): boolean => {
  let p1 = s.length,
    p2 = t.length;

  while (p1 >= 0 || p2 >= 0) {
    if (s[p1] === "#" || t[p2] === "#") {
      if (s[p1] === "#") {
        let backCount = 2;
        while (backCount > 0) {
          p1--;
          backCount--;

          if (s[p1] === "#") backCount += 2;
        }
      }
    }
    if (t[p2] === "#") {
      let backCount = 2;
      while (backCount > 0) {
        p2--;
        backCount--;
        if (t[p2] === "#") {
          backCount += 2;
        }
      }
    } else {
      if (s[p1] != t[p2]) return false;
      p1--;
      p2--;
    }
  }

  return true;
};

const lengthOfLongestSubstring = (s: string): number => {
  let longest = 0;
  let left = 0;
  const seenChars = new Map();

  for (let right = 0; right < s.length; right++) {
    const currentChar = s[right];
    const prevSeenChar = seenChars.get(currentChar);

    if (left <= prevSeenChar) {
      left = prevSeenChar + 1;
    }

    seenChars.set(currentChar, right);
    longest = Math.max(longest, right - left + 1);
  }

  return longest;
};

// console.log(backSpaceCompare("ab##", "c#d#"));
console.log(lengthOfLongestSubstring("abcabcbb"));
