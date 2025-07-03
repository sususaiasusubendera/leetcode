func kthCharacter(k int) byte {
    return byte('a' + bits.OnesCount(uint(k-1)))
}

// notice me senpai!